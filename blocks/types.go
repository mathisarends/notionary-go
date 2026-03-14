package blocks

import "strings"

type BlockType string

const (
	BlockTypeParagraph        BlockType = "paragraph"
	BlockTypeHeading1         BlockType = "heading_1"
	BlockTypeHeading2         BlockType = "heading_2"
	BlockTypeHeading3         BlockType = "heading_3"
	BlockTypeBulletedListItem BlockType = "bulleted_list_item"
	BlockTypeNumberedListItem BlockType = "numbered_list_item"
	BlockTypeCode             BlockType = "code"
	BlockTypeQuote            BlockType = "quote"
	BlockTypeCallout          BlockType = "callout"
	BlockTypeToDo             BlockType = "to_do"
	BlockTypeToggle           BlockType = "toggle"
	BlockTypeDivider          BlockType = "divider"
	BlockTypeImage            BlockType = "image"
	BlockTypeVideo            BlockType = "video"
	BlockTypeAudio            BlockType = "audio"
	BlockTypeFile             BlockType = "file"
	BlockTypePdf              BlockType = "pdf"
	BlockTypeBookmark         BlockType = "bookmark"
	BlockTypeEmbed            BlockType = "embed"
	BlockTypeEquation         BlockType = "equation"
	BlockTypeTable            BlockType = "table"
	BlockTypeTableRow         BlockType = "table_row"
	BlockTypeTableOfContents  BlockType = "table_of_contents"
	BlockTypeColumn           BlockType = "column"
	BlockTypeColumnList       BlockType = "column_list"
	BlockTypeChildPage        BlockType = "child_page"
	BlockTypeChildDatabase    BlockType = "child_database"
	BlockTypeSyncedBlock      BlockType = "synced_block"
	BlockTypeLinkPreview      BlockType = "link_preview"
	BlockTypeLinkToPage       BlockType = "link_to_page"
	BlockTypeBreadcrumb       BlockType = "breadcrumb"
	BlockTypeUnsupported      BlockType = "unsupported"
)

type BlockColor string

const (
	BlockColorDefault          BlockColor = "default"
	BlockColorBlue             BlockColor = "blue"
	BlockColorBrown            BlockColor = "brown"
	BlockColorGray             BlockColor = "gray"
	BlockColorGreen            BlockColor = "green"
	BlockColorOrange           BlockColor = "orange"
	BlockColorPink             BlockColor = "pink"
	BlockColorPurple           BlockColor = "purple"
	BlockColorRed              BlockColor = "red"
	BlockColorYellow           BlockColor = "yellow"
	BlockColorBlueBackground   BlockColor = "blue_background"
	BlockColorBrownBackground  BlockColor = "brown_background"
	BlockColorGrayBackground   BlockColor = "gray_background"
	BlockColorGreenBackground  BlockColor = "green_background"
	BlockColorOrangeBackground BlockColor = "orange_background"
	BlockColorPinkBackground   BlockColor = "pink_background"
	BlockColorPurpleBackground BlockColor = "purple_background"
	BlockColorRedBackground    BlockColor = "red_background"
	BlockColorYellowBackground BlockColor = "yellow_background"
)

func (c BlockColor) IsBackground() bool {
	return len(c) > 11 && c[len(c)-11:] == "_background"
}

func (c BlockColor) GetBaseColor() string {
	if c.IsBackground() {
		return string(c[:len(c)-11])
	}
	return string(c)
}

type CodeLanguage string

const (
	CodeLanguageAbap           CodeLanguage = "abap"
	CodeLanguageArduino        CodeLanguage = "arduino"
	CodeLanguageBash           CodeLanguage = "bash"
	CodeLanguageBasic          CodeLanguage = "basic"
	CodeLanguageC              CodeLanguage = "c"
	CodeLanguageClojure        CodeLanguage = "clojure"
	CodeLanguageCoffeescript   CodeLanguage = "coffeescript"
	CodeLanguageCpp            CodeLanguage = "c++"
	CodeLanguageCsharp         CodeLanguage = "c#"
	CodeLanguageCss            CodeLanguage = "css"
	CodeLanguageDart           CodeLanguage = "dart"
	CodeLanguageDiff           CodeLanguage = "diff"
	CodeLanguageDocker         CodeLanguage = "docker"
	CodeLanguageElixir         CodeLanguage = "elixir"
	CodeLanguageElm            CodeLanguage = "elm"
	CodeLanguageErlang         CodeLanguage = "erlang"
	CodeLanguageFlow           CodeLanguage = "flow"
	CodeLanguageFortran        CodeLanguage = "fortran"
	CodeLanguageFsharp         CodeLanguage = "f#"
	CodeLanguageGherkin        CodeLanguage = "gherkin"
	CodeLanguageGlsl           CodeLanguage = "glsl"
	CodeLanguageGo             CodeLanguage = "go"
	CodeLanguageGraphql        CodeLanguage = "graphql"
	CodeLanguageGroovy         CodeLanguage = "groovy"
	CodeLanguageHaskell        CodeLanguage = "haskell"
	CodeLanguageHtml           CodeLanguage = "html"
	CodeLanguageJava           CodeLanguage = "java"
	CodeLanguageJavascript     CodeLanguage = "javascript"
	CodeLanguageJson           CodeLanguage = "json"
	CodeLanguageJulia          CodeLanguage = "julia"
	CodeLanguageKotlin         CodeLanguage = "kotlin"
	CodeLanguageLatex          CodeLanguage = "latex"
	CodeLanguageLess           CodeLanguage = "less"
	CodeLanguageLisp           CodeLanguage = "lisp"
	CodeLanguageLivescript     CodeLanguage = "livescript"
	CodeLanguageLua            CodeLanguage = "lua"
	CodeLanguageMakefile       CodeLanguage = "makefile"
	CodeLanguageMarkdown       CodeLanguage = "markdown"
	CodeLanguageMarkup         CodeLanguage = "markup"
	CodeLanguageMatlab         CodeLanguage = "matlab"
	CodeLanguageMermaid        CodeLanguage = "mermaid"
	CodeLanguageNix            CodeLanguage = "nix"
	CodeLanguageObjectiveC     CodeLanguage = "objective-c"
	CodeLanguageOcaml          CodeLanguage = "ocaml"
	CodeLanguagePascal         CodeLanguage = "pascal"
	CodeLanguagePerl           CodeLanguage = "perl"
	CodeLanguagePhp            CodeLanguage = "php"
	CodeLanguagePlainText      CodeLanguage = "plain text"
	CodeLanguagePowershell     CodeLanguage = "powershell"
	CodeLanguageProlog         CodeLanguage = "prolog"
	CodeLanguageProtobuf       CodeLanguage = "protobuf"
	CodeLanguagePython         CodeLanguage = "python"
	CodeLanguageR              CodeLanguage = "r"
	CodeLanguageReason         CodeLanguage = "reason"
	CodeLanguageRuby           CodeLanguage = "ruby"
	CodeLanguageRust           CodeLanguage = "rust"
	CodeLanguageSass           CodeLanguage = "sass"
	CodeLanguageScala          CodeLanguage = "scala"
	CodeLanguageScheme         CodeLanguage = "scheme"
	CodeLanguageScss           CodeLanguage = "scss"
	CodeLanguageShell          CodeLanguage = "shell"
	CodeLanguageSql            CodeLanguage = "sql"
	CodeLanguageSwift          CodeLanguage = "swift"
	CodeLanguageTypescript     CodeLanguage = "typescript"
	CodeLanguageVbNet          CodeLanguage = "vb.net"
	CodeLanguageVerilog        CodeLanguage = "verilog"
	CodeLanguageVhdl           CodeLanguage = "vhdl"
	CodeLanguageVisualBasic    CodeLanguage = "visual basic"
	CodeLanguageWebassembly    CodeLanguage = "webassembly"
	CodeLanguageXml            CodeLanguage = "xml"
	CodeLanguageYaml           CodeLanguage = "yaml"
	CodeLanguageJavaCCppCsharp CodeLanguage = "java/c/c++/c#"
)

var CodeLanguageAliases = map[string]CodeLanguage{
	"cpp": CodeLanguageCpp,
	"c++": CodeLanguageCpp,
	"js":  CodeLanguageJavascript,
	"py":  CodeLanguagePython,
	"ts":  CodeLanguageTypescript,
}

var allCodeLanguages = []CodeLanguage{
	CodeLanguageAbap, CodeLanguageArduino, CodeLanguageBash, CodeLanguageBasic,
	CodeLanguageC, CodeLanguageClojure, CodeLanguageCoffeescript, CodeLanguageCpp,
	CodeLanguageCsharp, CodeLanguageCss, CodeLanguageDart, CodeLanguageDiff,
	CodeLanguageDocker, CodeLanguageElixir, CodeLanguageElm, CodeLanguageErlang,
	CodeLanguageFlow, CodeLanguageFortran, CodeLanguageFsharp, CodeLanguageGherkin,
	CodeLanguageGlsl, CodeLanguageGo, CodeLanguageGraphql, CodeLanguageGroovy,
	CodeLanguageHaskell, CodeLanguageHtml, CodeLanguageJava, CodeLanguageJavascript,
	CodeLanguageJson, CodeLanguageJulia, CodeLanguageKotlin, CodeLanguageLatex,
	CodeLanguageLess, CodeLanguageLisp, CodeLanguageLivescript, CodeLanguageLua,
	CodeLanguageMakefile, CodeLanguageMarkdown, CodeLanguageMarkup, CodeLanguageMatlab,
	CodeLanguageMermaid, CodeLanguageNix, CodeLanguageObjectiveC, CodeLanguageOcaml,
	CodeLanguagePascal, CodeLanguagePerl, CodeLanguagePhp, CodeLanguagePlainText,
	CodeLanguagePowershell, CodeLanguageProlog, CodeLanguageProtobuf, CodeLanguagePython,
	CodeLanguageR, CodeLanguageReason, CodeLanguageRuby, CodeLanguageRust,
	CodeLanguageSass, CodeLanguageScala, CodeLanguageScheme, CodeLanguageScss,
	CodeLanguageShell, CodeLanguageSql, CodeLanguageSwift, CodeLanguageTypescript,
	CodeLanguageVbNet, CodeLanguageVerilog, CodeLanguageVhdl, CodeLanguageVisualBasic,
	CodeLanguageWebassembly, CodeLanguageXml, CodeLanguageYaml, CodeLanguageJavaCCppCsharp,
}

func CodeLanguageFromString(lang string, defaultLang *CodeLanguage) CodeLanguage {
	if lang == "" {
		if defaultLang != nil {
			return *defaultLang
		}
		return CodeLanguagePlainText
	}

	normalized := strings.ToLower(strings.TrimSpace(lang))

	if alias, ok := CodeLanguageAliases[normalized]; ok {
		return alias
	}

	for _, l := range allCodeLanguages {
		if strings.ToLower(string(l)) == normalized {
			return l
		}
	}

	if defaultLang != nil {
		return *defaultLang
	}
	return CodeLanguagePlainText
}

type VideoFileType string

const (
	VideoFileTypeAmv  VideoFileType = ".amv"
	VideoFileTypeAsf  VideoFileType = ".asf"
	VideoFileTypeAvi  VideoFileType = ".avi"
	VideoFileTypeF4v  VideoFileType = ".f4v"
	VideoFileTypeFlv  VideoFileType = ".flv"
	VideoFileTypeGifv VideoFileType = ".gifv"
	VideoFileTypeMkv  VideoFileType = ".mkv"
	VideoFileTypeMov  VideoFileType = ".mov"
	VideoFileTypeMpg  VideoFileType = ".mpg"
	VideoFileTypeMpeg VideoFileType = ".mpeg"
	VideoFileTypeMpv  VideoFileType = ".mpv"
	VideoFileTypeMp4  VideoFileType = ".mp4"
	VideoFileTypeM4v  VideoFileType = ".m4v"
	VideoFileTypeQt   VideoFileType = ".qt"
	VideoFileTypeWmv  VideoFileType = ".wmv"
)

var allVideoFileTypes = []VideoFileType{
	VideoFileTypeAmv, VideoFileTypeAsf, VideoFileTypeAvi, VideoFileTypeF4v,
	VideoFileTypeFlv, VideoFileTypeGifv, VideoFileTypeMkv, VideoFileTypeMov,
	VideoFileTypeMpg, VideoFileTypeMpeg, VideoFileTypeMpv, VideoFileTypeMp4,
	VideoFileTypeM4v, VideoFileTypeQt, VideoFileTypeWmv,
}

func VideoFileTypeGetAllExtensions() map[VideoFileType]struct{} {
	result := make(map[VideoFileType]struct{}, len(allVideoFileTypes))
	for _, ext := range allVideoFileTypes {
		result[ext] = struct{}{}
	}
	return result
}

func VideoFileTypeIsValidExtension(filename string) bool {
	lower := strings.ToLower(filename)
	for _, ext := range allVideoFileTypes {
		if strings.HasSuffix(lower, string(ext)) {
			return true
		}
	}
	return false
}