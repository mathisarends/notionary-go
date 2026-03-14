package elements

import (
	"strings"
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func richTextPlain(rts []blocks.RichText) string {
	var b strings.Builder
	for _, rt := range rts {
		b.WriteString(rt.PlainText)
	}
	return b.String()
}

func expectParse[T blocks.Block](t *testing.T, codec BlockCodec, line string) T {
	t.Helper()
	block, ok := codec.Parse(line)
	if !ok {
		t.Fatalf("expected parse success for %q", line)
	}
	typed, ok := block.(T)
	if !ok {
		t.Fatalf("unexpected block type %T", block)
	}
	return typed
}

func expectParseFail(t *testing.T, codec BlockCodec, line string) {
	t.Helper()
	if block, ok := codec.Parse(line); ok {
		t.Fatalf("expected parse failure for %q, got %T", line, block)
	}
}

func expectRender(t *testing.T, codec BlockCodec, block blocks.Block, want string) {
	t.Helper()
	got, ok := codec.Render(block)
	if !ok {
		t.Fatalf("expected render success for %T", block)
	}
	if got != want {
		t.Fatalf("unexpected render output. want %q, got %q", want, got)
	}
}

func expectRenderFail(t *testing.T, codec BlockCodec, block blocks.Block) {
	t.Helper()
	if got, ok := codec.Render(block); ok {
		t.Fatalf("expected render failure for %T, got %q", block, got)
	}
}

func TestAudioCodec(t *testing.T) {
	codec := &AudioCodec{}
	block := expectParse[*blocks.AudioBlock](t, codec, `<audio src="https://example.com/track.mp3" caption="clip">`)
	if block.Audio.URL != "https://example.com/track.mp3" {
		t.Fatalf("unexpected audio url %q", block.Audio.URL)
	}
	if got := richTextPlain(block.Audio.Caption); got != "clip" {
		t.Fatalf("unexpected caption %q", got)
	}
	expectParseFail(t, codec, `<audio src=>`)
	expectRender(t, codec, block, `<audio src="https://example.com/track.mp3" caption="clip">`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestBookmarkCodec(t *testing.T) {
	codec := &BookmarkCodec{}
	block := expectParse[*blocks.BookmarkBlock](t, codec, `<bookmark url="https://example.com" title="Example">`)
	if block.Bookmark.URL != "https://example.com" {
		t.Fatalf("unexpected bookmark url %q", block.Bookmark.URL)
	}
	if got := richTextPlain(block.Bookmark.Title); got != "Example" {
		t.Fatalf("unexpected title %q", got)
	}
	expectParseFail(t, codec, `<bookmark title="Missing URL">`)
	expectRender(t, codec, block, `<bookmark url="https://example.com" title="Example">`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestBreadcrumbCodec(t *testing.T) {
	codec := &BreadcrumbCodec{}
	_ = expectParse[*blocks.BreadcrumbBlock](t, codec, `[breadcrumb]`)
	expectParseFail(t, codec, `[bread]`)
	expectRender(t, codec, &blocks.BreadcrumbBlock{}, `[breadcrumb]`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestBulletedListCodec(t *testing.T) {
	codec := &BulletedListCodec{}
	block := expectParse[*blocks.BulletedListItemBlock](t, codec, `- item`)
	if got := richTextPlain(block.BulletedListItem.RichText); got != "item" {
		t.Fatalf("unexpected list text %q", got)
	}
	expectParseFail(t, codec, `- [ ] todo`)
	expectRender(t, codec, block, `- item`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestCalloutCodec(t *testing.T) {
	codec := &CalloutCodec{}
	block := expectParse[*blocks.CalloutBlock](t, codec, `<callout emoji="!" color="blue">`)
	if block.Callout.Emoji != "!" {
		t.Fatalf("unexpected emoji %q", block.Callout.Emoji)
	}
	if block.Callout.Color != blocks.BlockColorBlue {
		t.Fatalf("unexpected color %q", block.Callout.Color)
	}
	expectParseFail(t, codec, `<callout wrong>`)
	expectRender(t, codec, block, `<callout emoji="!" color="blue">`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestCodeCodec(t *testing.T) {
	codec := &CodeCodec{}
	block := expectParse[*blocks.CodeBlock](t, codec, `<code lang="go">`)
	if block.Code.Language != "go" {
		t.Fatalf("unexpected language %q", block.Code.Language)
	}
	expectParseFail(t, codec, `<code language="go">`)
	expectRender(t, codec, block, `<code lang="go">`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestColumnListCodec(t *testing.T) {
	codec := &ColumnListCodec{}
	_ = expectParse[*blocks.ColumnListBlock](t, codec, `<columns>`)
	expectParseFail(t, codec, `<column>`)
	expectRender(t, codec, &blocks.ColumnListBlock{}, `<columns>>`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestColumnCodec(t *testing.T) {
	codec := &ColumnCodec{}
	block := expectParse[*blocks.ColumnBlock](t, codec, `<column ratio="0.5">`)
	if block.Column.Ratio != "0.5" {
		t.Fatalf("unexpected ratio %q", block.Column.Ratio)
	}
	expectParseFail(t, codec, `<column ratio="2">`)
	expectRender(t, codec, block, `<column ratio="0.5">`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestDividerCodec(t *testing.T) {
	codec := &DividerCodec{}
	_ = expectParse[*blocks.DividerBlock](t, codec, `---`)
	expectParseFail(t, codec, `--`)
	expectRender(t, codec, &blocks.DividerBlock{}, `---`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestEmbedCodec(t *testing.T) {
	codec := &EmbedCodec{}
	block := expectParse[*blocks.EmbedBlock](t, codec, `<embed url="https://example.com/embed" title="Preview">`)
	if block.Embed.URL != "https://example.com/embed" {
		t.Fatalf("unexpected embed url %q", block.Embed.URL)
	}
	if got := richTextPlain(block.Embed.Title); got != "Preview" {
		t.Fatalf("unexpected title %q", got)
	}
	expectParseFail(t, codec, `<embed url="notaurl">`)
	expectRender(t, codec, block, `<embed url="https://example.com/embed" title="Preview">`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestEquationCodec(t *testing.T) {
	codec := &EquationCodec{}
	_ = expectParse[*blocks.EquationBlock](t, codec, `<equation>`)
	expectParseFail(t, codec, `<equation a="b">`)
	expectRender(t, codec, &blocks.EquationBlock{}, `<equation>>`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestFileCodec(t *testing.T) {
	codec := &FileCodec{}
	block := expectParse[*blocks.FileBlock](t, codec, `<file src="https://example.com/a.zip" name="archive">`)
	if block.File.URL != "https://example.com/a.zip" {
		t.Fatalf("unexpected file url %q", block.File.URL)
	}
	if block.File.Name != "archive" {
		t.Fatalf("unexpected file name %q", block.File.Name)
	}
	expectParseFail(t, codec, `<file name="missing">`)
	expectRender(t, codec, block, `<file src="https://example.com/a.zip" name="archive">`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestHeadingCodec(t *testing.T) {
	codec := &HeadingCodec{}
	h1 := expectParse[*blocks.Heading1Block](t, codec, `# A`)
	if got := richTextPlain(h1.Heading1.RichText); got != "A" {
		t.Fatalf("unexpected h1 text %q", got)
	}
	h2 := expectParse[*blocks.Heading2Block](t, codec, `## B`)
	if got := richTextPlain(h2.Heading2.RichText); got != "B" {
		t.Fatalf("unexpected h2 text %q", got)
	}
	h3 := expectParse[*blocks.Heading3Block](t, codec, `### C`)
	if got := richTextPlain(h3.Heading3.RichText); got != "C" {
		t.Fatalf("unexpected h3 text %q", got)
	}
	expectParseFail(t, codec, `#### Too deep`)
	expectRender(t, codec, h1, `# A`)
	expectRender(t, codec, h2, `## B`)
	expectRender(t, codec, h3, `### C`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestImageCodec(t *testing.T) {
	codec := &ImageCodec{}
	block := expectParse[*blocks.ImageBlock](t, codec, `<image src="https://example.com/photo.png" caption="alt">`)
	if block.Image.URL != "https://example.com/photo.png" {
		t.Fatalf("unexpected image url %q", block.Image.URL)
	}
	if got := richTextPlain(block.Image.Caption); got != "alt" {
		t.Fatalf("unexpected caption %q", got)
	}
	expectParseFail(t, codec, `<image caption="missing src">`)
	expectRender(t, codec, block, `<image src="https://example.com/photo.png" caption="alt">`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestNumberedListCodec(t *testing.T) {
	codec := &NumberedListCodec{}
	block := expectParse[*blocks.NumberedListItemBlock](t, codec, `42. step`)
	if got := richTextPlain(block.NumberedListItem.RichText); got != "step" {
		t.Fatalf("unexpected list text %q", got)
	}
	expectParseFail(t, codec, `x. step`)
	expectRender(t, codec, block, `1. step`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestParagraphCodec(t *testing.T) {
	codec := &ParagraphCodec{}
	block := expectParse[*blocks.ParagraphBlock](t, codec, `plain text`)
	if got := richTextPlain(block.Paragraph.RichText); got != "plain text" {
		t.Fatalf("unexpected paragraph text %q", got)
	}
	expectRender(t, codec, block, `plain text`)
	expectRenderFail(t, codec, &blocks.Heading1Block{})
}

func TestPDFCodec(t *testing.T) {
	codec := &PDFCodec{}
	block := expectParse[*blocks.PDFBlock](t, codec, `<pdf src="https://example.com/doc.pdf" caption="docs">`)
	if block.PDF.URL != "https://example.com/doc.pdf" {
		t.Fatalf("unexpected pdf url %q", block.PDF.URL)
	}
	if got := richTextPlain(block.PDF.Caption); got != "docs" {
		t.Fatalf("unexpected caption %q", got)
	}
	expectParseFail(t, codec, `<pdf caption="missing src">`)
	expectRender(t, codec, block, `<pdf src="https://example.com/doc.pdf" caption="docs">`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestQuoteCodec(t *testing.T) {
	codec := &QuoteCodec{}
	block := expectParse[*blocks.QuoteBlock](t, codec, `> cited`)
	if got := richTextPlain(block.Quote.RichText); got != "cited" {
		t.Fatalf("unexpected quote text %q", got)
	}
	expectParseFail(t, codec, `>> not a quote`)
	expectRender(t, codec, block, `> cited`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestSyncedBlockCodec(t *testing.T) {
	codec := &SyncedBlockCodec{}
	block := expectParse[*blocks.SyncedBlock](t, codec, `<synced id="abc">`)
	if block.SyncedBlock.ID != "abc" {
		t.Fatalf("unexpected synced id %q", block.SyncedBlock.ID)
	}
	expectParseFail(t, codec, `<sync id="abc">`)
	expectRender(t, codec, block, `<synced id="abc">`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestTableOfContentsCodec(t *testing.T) {
	codec := &TableOfContentsCodec{}
	_ = expectParse[*blocks.TableOfContentsBlock](t, codec, `[toc]`)
	expectParseFail(t, codec, `[to]`)
	expectRender(t, codec, &blocks.TableOfContentsBlock{}, `[toc]`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestTableCodec(t *testing.T) {
	codec := &TableCodec{}
	_ = expectParse[*blocks.TableBlock](t, codec, `| A | B |`)
	expectParseFail(t, codec, `A | B`)
	expectRender(t, codec, &blocks.TableBlock{}, `|`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestTableRowCodec(t *testing.T) {
	codec := &TableRowCodec{}
	_ = expectParse[*blocks.TableRowBlock](t, codec, `| --- | --- |`)
	expectParseFail(t, codec, `| A | B |`)
	expectRender(t, codec, &blocks.TableRowBlock{}, `|`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestToDoCodec(t *testing.T) {
	codec := &ToDoCodec{}
	block := expectParse[*blocks.ToDoBlock](t, codec, `- [ ] todo`)
	if block.ToDo.Checked {
		t.Fatalf("expected unchecked todo")
	}
	if got := richTextPlain(block.ToDo.RichText); got != "todo" {
		t.Fatalf("unexpected todo text %q", got)
	}
	expectParseFail(t, codec, `- [x] done`)
	expectRender(t, codec, block, `- [ ] todo`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
	expectRenderFail(t, codec, &blocks.ToDoBlock{ToDo: blocks.ToDoData{Checked: true}})
}

func TestToDoDoneCodec(t *testing.T) {
	codec := &ToDoDoneCodec{}
	block := expectParse[*blocks.ToDoBlock](t, codec, `- [x] done`)
	if !block.ToDo.Checked {
		t.Fatalf("expected checked todo")
	}
	if got := richTextPlain(block.ToDo.RichText); got != "done" {
		t.Fatalf("unexpected todo text %q", got)
	}
	expectParseFail(t, codec, `- [ ] todo`)
	expectRender(t, codec, block, `- [x] done`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
	expectRenderFail(t, codec, &blocks.ToDoBlock{ToDo: blocks.ToDoData{Checked: false}})
}

func TestToggleCodec(t *testing.T) {
	codec := &ToggleCodec{}
	block := expectParse[*blocks.ToggleBlock](t, codec, `<toggle title="Section">`)
	if got := richTextPlain(block.Toggle.Title); got != "Section" {
		t.Fatalf("unexpected toggle title %q", got)
	}
	expectParseFail(t, codec, `<toggle heading="Section">`)
	expectRender(t, codec, block, `<toggle title="Section">`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestToggleableHeadingCodec(t *testing.T) {
	codec := &ToggleableHeadingCodec{}
	block := expectParse[*blocks.ToggleableHeadingBlock](t, codec, `<toggle title="Section" level="2">`)
	if got := richTextPlain(block.ToggleableHeading.Title); got != "Section" {
		t.Fatalf("unexpected heading title %q", got)
	}
	if block.ToggleableHeading.Level != 2 {
		t.Fatalf("unexpected heading level %d", block.ToggleableHeading.Level)
	}
	expectParseFail(t, codec, `<toggle title="Section" level="4">`)
	expectRender(t, codec, block, `<toggle title="Section" level="2">`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}

func TestVideoCodec(t *testing.T) {
	codec := &VideoCodec{}
	block := expectParse[*blocks.VideoBlock](t, codec, `<video src="https://example.com/clip.mp4" caption="demo">`)
	if block.Video.URL != "https://example.com/clip.mp4" {
		t.Fatalf("unexpected video url %q", block.Video.URL)
	}
	if got := richTextPlain(block.Video.Caption); got != "demo" {
		t.Fatalf("unexpected caption %q", got)
	}
	expectParseFail(t, codec, `<video caption="missing src">`)
	expectRender(t, codec, block, `<video src="https://example.com/clip.mp4" caption="demo">`)
	expectRenderFail(t, codec, &blocks.ParagraphBlock{})
}
