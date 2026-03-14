package markdown

import "regexp"

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
}

// SelfClosingTagSyntax represents a single-line HTML-style tag with no closing counterpart.
// Used for media elements like <image src="...">, <audio src="...">, etc.
type SelfClosingTagSyntax struct {
	Tag     string        
	Pattern *regexp.Regexp 
}

// TagSyntax represents a multi-line HTML-style block with an opening and closing tag.
// Used for containers like <callout>, <toggle>, <code>, etc.
type TagSyntax struct {
	OpenTag    string        
	CloseTag   string         
	Pattern    *regexp.Regexp 
	EndPattern *regexp.Regexp 
}

func (SimpleSyntax) isSyntaxDefinition()        {}
func (SelfClosingTagSyntax) isSyntaxDefinition() {}
func (TagSyntax) isSyntaxDefinition()            {}

var Registry = map[RegistryKey]Definition{
	Audio:    SelfClosingTagSyntax{Tag: "<audio",    Pattern: regexp.MustCompile(`^<audio\s+src="([^"]+)"(?:\s+caption="([^"]*)")?>\s*$`)},
	Bookmark: SelfClosingTagSyntax{Tag: "<bookmark", Pattern: regexp.MustCompile(`^<bookmark\s+url="(https?://[^\s"]+)"(?:\s+title="([^"]*)")?>\s*$`)},
	Image:    SelfClosingTagSyntax{Tag: "<image",    Pattern: regexp.MustCompile(`^<image\s+src="([^"]+)"(?:\s+caption="([^"]*)")?>\s*$`)},
	Video:    SelfClosingTagSyntax{Tag: "<video",    Pattern: regexp.MustCompile(`^<video\s+src="([^"]+)"(?:\s+caption="([^"]*)")?>\s*$`)},
	File:     SelfClosingTagSyntax{Tag: "<file",     Pattern: regexp.MustCompile(`^<file\s+src="([^"]+)"(?:\s+name="([^"]*)")?>\s*$`)},
	PDF:      SelfClosingTagSyntax{Tag: "<pdf",      Pattern: regexp.MustCompile(`^<pdf\s+src="([^"]+)"(?:\s+caption="([^"]*)")?>\s*$`)},
	Embed:    SelfClosingTagSyntax{Tag: "<embed",    Pattern: regexp.MustCompile(`^<embed\s+url="(https?://[^\s"]+)"(?:\s+title="([^"]*)")?>\s*$`)},

	BulletedList: SimpleSyntax{StartDelimiter: "- ",    Pattern: regexp.MustCompile(`^(\s*)-\s+(?!\[[ xX]\])(.+)$`)},
	NumberedList: SimpleSyntax{StartDelimiter: "1. ",   Pattern: regexp.MustCompile(`^(\s*)(\d+)\.\s+(.+)$`)},
	ToDo:         SimpleSyntax{StartDelimiter: "- [ ]", Pattern: regexp.MustCompile(`^\s*-\s+\[ \]\s+(.+)$`)},
	ToDoDone:     SimpleSyntax{StartDelimiter: "- [x]", Pattern: regexp.MustCompile(`(?i)^\s*-\s+\[x\]\s+(.+)$`)},

	Callout: TagSyntax{
		OpenTag:    "<callout",
		CloseTag:   "</callout>",
		Pattern:    regexp.MustCompile(`^<callout(?:\s+emoji="([^"]*)")?(?:\s+color="([^"]*)")?>\s*$`),
		EndPattern: regexp.MustCompile(`^</callout>\s*$`),
	},
	Toggle: TagSyntax{
		OpenTag:    "<toggle",
		CloseTag:   "</toggle>",
		Pattern:    regexp.MustCompile(`^<toggle(?:\s+title="([^"]*)")?>\s*$`),
		EndPattern: regexp.MustCompile(`^</toggle>\s*$`),
	},
	ToggleableHeading: TagSyntax{
		OpenTag:    "<toggle",
		CloseTag:   "</toggle>",
		Pattern:    regexp.MustCompile(`^<toggle\s+title="([^"]*)"\s+level="([123])">\s*$`),
		EndPattern: regexp.MustCompile(`^</toggle>\s*$`),
	},
	Code: TagSyntax{
		OpenTag:    "<code",
		CloseTag:   "</code>",
		Pattern:    regexp.MustCompile(`^<code(?:\s+lang="([^"]*)")?>\s*$`),
		EndPattern: regexp.MustCompile(`^</code>\s*$`),
	},
	Equation: TagSyntax{
		OpenTag:    "<equation>",
		CloseTag:   "</equation>",
		Pattern:    regexp.MustCompile(`^<equation>\s*$`),
		EndPattern: regexp.MustCompile(`^</equation>\s*$`),
	},
	ColumnList: TagSyntax{
		OpenTag:    "<columns>",
		CloseTag:   "</columns>",
		Pattern:    regexp.MustCompile(`^<columns>\s*$`),
		EndPattern: regexp.MustCompile(`^</columns>\s*$`),
	},
	Column: TagSyntax{
		OpenTag:    "<column",
		CloseTag:   "</column>",
		Pattern:    regexp.MustCompile(`^<column(?:\s+ratio="(0?\.\d+|1(?:\.0?)?)")?>\s*$`),
		EndPattern: regexp.MustCompile(`^</column>\s*$`),
	},
	SyncedBlock: TagSyntax{
		OpenTag:    "<synced",
		CloseTag:   "</synced>",
		Pattern:    regexp.MustCompile(`^<synced(?:\s+id="([^"]*)")?>\s*$`),
		EndPattern: regexp.MustCompile(`^</synced>\s*$`),
	},

	Heading:         SimpleSyntax{StartDelimiter: "#",           Pattern: regexp.MustCompile(`^(#{1,3})[ \t]+(.+)$`)},
	Quote:           SimpleSyntax{StartDelimiter: "> ",          Pattern: regexp.MustCompile(`^>(?!>)\s*(.+)$`)},
	Divider:         SimpleSyntax{StartDelimiter: "---",         Pattern: regexp.MustCompile(`^\s*-{3,}\s*$`)},
	Breadcrumb:      SimpleSyntax{StartDelimiter: "[breadcrumb]",Pattern: regexp.MustCompile(`(?i)^\[breadcrumb\]\s*$`)},
	TableOfContents: SimpleSyntax{StartDelimiter: "[toc]",       Pattern: regexp.MustCompile(`(?i)^\[toc\]$`)},
	Table:           SimpleSyntax{StartDelimiter: "|",           Pattern: regexp.MustCompile(`^\s*\|(.+)\|\s*$`)},
	TableRow:        SimpleSyntax{StartDelimiter: "|",           Pattern: regexp.MustCompile(`^\s*\|([\s\-:|]+)\|\s*$`)},
	Caption:         SimpleSyntax{StartDelimiter: "[caption]",   Pattern: regexp.MustCompile(`^\[caption\]\s+(\S.*)$`)},
	Space:           SimpleSyntax{StartDelimiter: "[space]",     Pattern: regexp.MustCompile(`^\[space\]\s*$`)},
}