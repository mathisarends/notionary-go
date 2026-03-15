package markdown

import (
	"regexp"
	"strings"
)

type RegistryKey string

const (
	Audio    RegistryKey = "audio"
	Bookmark RegistryKey = "bookmark"
	Image    RegistryKey = "image"
	Video    RegistryKey = "video"
	File     RegistryKey = "file"
	PDF      RegistryKey = "pdf"
	Embed    RegistryKey = "embed"

	BulletedList RegistryKey = "bulleted_list"
	NumberedList RegistryKey = "numbered_list"
	ToDo         RegistryKey = "todo"
	ToDoDone     RegistryKey = "todo_done"

	Toggle            RegistryKey = "toggle"
	ToggleableHeading RegistryKey = "toggleable_heading"
	Callout           RegistryKey = "callout"
	Code              RegistryKey = "code"
	SyncedBlock       RegistryKey = "synced_block"
	Equation          RegistryKey = "equation"
	ColumnList        RegistryKey = "column_list"
	Column            RegistryKey = "column"

	Quote           RegistryKey = "quote"
	Heading         RegistryKey = "heading"
	Divider         RegistryKey = "divider"
	Breadcrumb      RegistryKey = "breadcrumb"
	TableOfContents RegistryKey = "table_of_contents"
	Table           RegistryKey = "table"
	TableRow        RegistryKey = "table_row"
	Caption         RegistryKey = "caption"
	Space           RegistryKey = "space"
	Paragraph       RegistryKey = "paragraph"
)

type Definition interface{ isSyntaxDefinition() }

// SimpleSyntax matches a single line via a delimiter prefix and regex pattern.
// Used for headings, lists, quotes, dividers, etc.
type SimpleSyntax struct {
	StartDelimiter string
	Pattern        *regexp.Regexp
	Group          string
	Example        string
}

// SelfClosingTagSyntax represents a single-line HTML-style tag with no closing counterpart.
// Used for media elements like <image src="...">, <audio src="...">, etc.
type SelfClosingTagSyntax struct {
	Tag     string        
	Pattern *regexp.Regexp 
	Group   string
	Example string
}

// TagSyntax represents a multi-line HTML-style block with an opening and closing tag.
// Used for containers like <callout>, <toggle>, <code>, etc.
type TagSyntax struct {
	OpenTag    string        
	CloseTag   string         
	Pattern    *regexp.Regexp 
	EndPattern *regexp.Regexp
	Children   bool
	Group      string
	Example    string
}

func (SimpleSyntax) isSyntaxDefinition()        {}
func (SelfClosingTagSyntax) isSyntaxDefinition() {}
func (TagSyntax) isSyntaxDefinition()            {}

var Registry = map[RegistryKey]Definition{
	Audio:    SelfClosingTagSyntax{Tag: "<audio", Pattern: regexp.MustCompile(`^<audio\s+src="([^"]+)"(?:\s+caption="([^"]*)")?>\s*$`), Group: "Media", Example: `<audio src="https://example.com/track.mp3" caption="optional caption">`},
	Bookmark: SelfClosingTagSyntax{Tag: "<bookmark", Pattern: regexp.MustCompile(`^<bookmark\s+url="(https?://[^\s"]+)"(?:\s+title="([^"]*)")?>\s*$`), Group: "Links", Example: `<bookmark url="https://example.com" title="optional title">`},
	Image:    SelfClosingTagSyntax{Tag: "<image", Pattern: regexp.MustCompile(`^<image\s+src="([^"]+)"(?:\s+caption="([^"]*)")?>\s*$`), Group: "Media", Example: `<image src="https://example.com/photo.png" caption="optional caption">`},
	Video:    SelfClosingTagSyntax{Tag: "<video", Pattern: regexp.MustCompile(`^<video\s+src="([^"]+)"(?:\s+caption="([^"]*)")?>\s*$`), Group: "Media", Example: `<video src="https://example.com/clip.mp4" caption="optional caption">`},
	File:     SelfClosingTagSyntax{Tag: "<file", Pattern: regexp.MustCompile(`^<file\s+src="([^"]+)"(?:\s+name="([^"]*)")?>\s*$`), Group: "Media", Example: `<file src="https://example.com/data.zip" name="optional display name">`},
	PDF:      SelfClosingTagSyntax{Tag: "<pdf", Pattern: regexp.MustCompile(`^<pdf\s+src="([^"]+)"(?:\s+caption="([^"]*)")?>\s*$`), Group: "Media", Example: `<pdf src="https://example.com/doc.pdf" caption="optional caption">`},
	Embed:    SelfClosingTagSyntax{Tag: "<embed", Pattern: regexp.MustCompile(`^<embed\s+url="(https?://[^\s"]+)"(?:\s+title="([^"]*)")?>\s*$`), Group: "Links", Example: `<embed url="https://example.com/embed" title="optional title">`},

	BulletedList: SimpleSyntax{StartDelimiter: "- ", Pattern: regexp.MustCompile(`^(\s*)-\s+(.+)$`), Group: "Lists", Example: "- First item\n- Second item\n  - Nested item"},
	NumberedList: SimpleSyntax{StartDelimiter: "1. ", Pattern: regexp.MustCompile(`^(\s*)(\d+)\.\s+(.+)$`), Group: "Lists", Example: "1. First item\n2. Second item"},
	ToDo:         SimpleSyntax{StartDelimiter: "- [ ]", Pattern: regexp.MustCompile(`^\s*-\s+\[ \]\s+(.+)$`), Group: "Lists", Example: "- [ ] Task to do"},
	ToDoDone:     SimpleSyntax{StartDelimiter: "- [x]", Pattern: regexp.MustCompile(`(?i)^\s*-\s+\[x\]\s+(.+)$`), Group: "Lists", Example: "- [x] Completed task"},

	Heading:         SimpleSyntax{StartDelimiter: "#", Pattern: regexp.MustCompile(`^(#{1,3})[ \t]+(.+)$`), Group: "Text", Example: "# Heading 1\n## Heading 2\n### Heading 3"},
	Quote:           SimpleSyntax{StartDelimiter: "> ", Pattern: regexp.MustCompile(`^>\s*(.+)$`), Group: "Text", Example: "> This is a blockquote."},
	Divider:         SimpleSyntax{StartDelimiter: "---", Pattern: regexp.MustCompile(`^\s*-{3,}\s*$`), Group: "Text", Example: "---"},
	Breadcrumb:      SimpleSyntax{StartDelimiter: "[breadcrumb]", Pattern: regexp.MustCompile(`(?i)^\[breadcrumb\]\s*$`), Group: "Text", Example: "[breadcrumb]"},
	TableOfContents: SimpleSyntax{StartDelimiter: "[toc]", Pattern: regexp.MustCompile(`(?i)^\[toc\]$`), Group: "Text", Example: "[toc]"},
	Caption:         SimpleSyntax{StartDelimiter: "[caption]", Pattern: regexp.MustCompile(`^\[caption\]\s+(\S.*)$`), Group: "Text", Example: "[caption] This describes the block above."},
	Space:           SimpleSyntax{StartDelimiter: "[space]", Pattern: regexp.MustCompile(`^\[space\]\s*$`), Group: "Text", Example: "[space]"},
	Paragraph:       SimpleSyntax{StartDelimiter: "", Pattern: nil, Group: "Text", Example: "Any plain text line becomes a paragraph."},

	Table:    SimpleSyntax{StartDelimiter: "|", Pattern: regexp.MustCompile(`^\s*\|(.+)\|\s*$`), Group: "Table", Example: "| Col A | Col B |\n|-------|-------|\n| val 1 | val 2 |"},
	TableRow: SimpleSyntax{StartDelimiter: "|", Pattern: regexp.MustCompile(`^\s*\|([\s\-:|]+)\|\s*$`), Group: "Table", Example: "|-------|-------|"},

	Callout: TagSyntax{
		OpenTag: "<callout", CloseTag: "</callout>",
		Pattern: regexp.MustCompile(`^<callout(?:\s+emoji="([^"]*)")?(?:\s+color="([^"]*)")?>\s*$`), EndPattern: regexp.MustCompile(`^</callout>\s*$`),
		Children: true, Group: "Containers",
		Example: "<callout emoji=\"💡\" color=\"blue\">\nContent here — supports nested blocks.\n</callout>",
	},
	Toggle: TagSyntax{
		OpenTag: "<toggle", CloseTag: "</toggle>",
		Pattern: regexp.MustCompile(`^<toggle(?:\s+title="([^"]*)")?>\s*$`), EndPattern: regexp.MustCompile(`^</toggle>\s*$`),
		Children: true, Group: "Containers",
		Example: "<toggle title=\"Click to expand\">\nHidden content — supports nested blocks.\n</toggle>",
	},
	ToggleableHeading: TagSyntax{
		OpenTag: "<toggle", CloseTag: "</toggle>",
		Pattern: regexp.MustCompile(`^<toggle\s+title="([^"]*)"\s+level="([123])">\s*$`), EndPattern: regexp.MustCompile(`^</toggle>\s*$`),
		Children: true, Group: "Containers",
		Example: "<toggle title=\"Section Title\" level=\"2\">\nContent under the heading toggle.\n</toggle>",
	},
	Code: TagSyntax{
		OpenTag: "```", CloseTag: "```",
		Pattern: regexp.MustCompile("^\\s*```([^`\\s]*)\\s*$"), EndPattern: regexp.MustCompile("^\\s*```\\s*$"),
		Group: "Containers", Example: "```python\nprint(\"hello\")\n```",
	},
	Equation: TagSyntax{
		OpenTag: "<equation>", CloseTag: "</equation>",
		Pattern: regexp.MustCompile(`^<equation>\s*$`), EndPattern: regexp.MustCompile(`^</equation>\s*$`),
		Group: "Containers", Example: "<equation>\nE = mc^2\n</equation>",
	},
	ColumnList: TagSyntax{
		OpenTag: "<columns>", CloseTag: "</columns>",
		Pattern: regexp.MustCompile(`^<columns>\s*$`), EndPattern: regexp.MustCompile(`^</columns>\s*$`),
		Children: true, Group: "Containers",
		Example: "<columns>\n<column ratio=\"0.5\">\nLeft content.\n</column>\n<column ratio=\"0.5\">\nRight content.\n</column>\n</columns>",
	},
	Column: TagSyntax{
		OpenTag: "<column", CloseTag: "</column>",
		Pattern: regexp.MustCompile(`^<column(?:\s+ratio="(0?\.\d+|1(?:\.0?)?)")?>\s*$`), EndPattern: regexp.MustCompile(`^</column>\s*$`),
		Children: true, Group: "Containers",
		Example: "<column ratio=\"0.5\">\nColumn content — must be a child of <columns>.\n</column>",
	},
	SyncedBlock: TagSyntax{
		OpenTag: "<synced", CloseTag: "</synced>",
		Pattern: regexp.MustCompile(`^<synced(?:\s+id="([^"]*)")?>\s*$`), EndPattern: regexp.MustCompile(`^</synced>\s*$`),
		Children: true, Group: "Containers",
		Example: "<synced id=\"optional-id\">\nContent that can be reused across pages.\n</synced>",
	},
}

var GuideOrder = []RegistryKey{
	Paragraph, Heading, Quote, Divider, Space, Breadcrumb, TableOfContents, Caption,
	BulletedList, NumberedList, ToDo, ToDoDone,
	Table, TableRow,
	Image, Video, Audio, PDF, File,
	Bookmark, Embed,
	Callout, Toggle, ToggleableHeading, Code, Equation, ColumnList, Column, SyncedBlock,
}

// SyntaxGuide generates a human-readable syntax reference from the Registry.
// The output is suitable for use in agent system prompts.
func SyntaxGuide() string {
	var b strings.Builder
	b.WriteString("# Notionary Markdown Syntax Reference\n\n")

	lastGroup := ""
	for _, key := range GuideOrder {
		def, ok := Registry[key]
		if !ok {
			continue
		}

		group, example := extractMeta(def)
		if example == "" {
			continue
		}

		if group != lastGroup {
			b.WriteString("## " + group + "\n\n")
			lastGroup = group
		}

		b.WriteString("### " + string(key) + "\n")
		b.WriteString("````\n")
		b.WriteString(example)
		b.WriteString("\n````\n\n")
	}

	return b.String()
}

func extractMeta(def Definition) (group, example string) {
	switch d := def.(type) {
	case SimpleSyntax:
		return d.Group, d.Example
	case SelfClosingTagSyntax:
		return d.Group, d.Example
	case TagSyntax:
		return d.Group, d.Example
	}
	return "", ""
}