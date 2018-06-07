package pretty

const githubURL = "https://github.com/vonEdfa/go-pretty-strings"
const bugReportURL = "https://github.com/vonEdfa/go-pretty-strings/issues"

/**
 * NOTE: If you update a default, please update the comment in the appropriate struct too in structs.go!
 */

const (
	defaultAutomaticMaxWidth = true
	defaultBreakpoints       = " -"
	defaultNewline           = "\n"
	defaultParagraph         = "\n\n"

	// Border defaults
	defaultBorderColor          = "37" //White
	defaultBorderTitleBold      = false
	defaultBorderTitleCase      = 2
	defaultBorderTitleColor     = defaultBorderColor
	defaultBorderTitleItalics   = false
	defaultBorderTitleSpaced    = false
	defaultBorderTitleUnderline = false
	defaultBorderUseTitleColor  = true

	// OutputMargin defaults
	defaultMarginChar  = "\u00A0"
	defaultMarginWidth = 2

	// OutputPadding defaults
	defaultPaddingChar  = "\u00A0"
	defaultPaddingWidth = 3

	// Titles defaults
	defaultTitleBold      = true
	defaultTitleCase      = TitleCase
	defaultTitleColor     = "37" //White
	defaultTitleItalics   = false
	defaultTitleSpaced    = false
	defaultTitleUnderline = false
	defaultUseTitleColor  = true

	// WordWrap defaults
	defaultBreakWord          = true
	defaultHyphen             = "-"
	defaultUseTrailingNewLine = false
)

// Cases
const (
	InheritCase int = iota
	UpperCase
	TitleCase
	LowerCase
)

// New returns a new instance of Pretty, initialized with the default settings.
func New() Pretty {
	return Pretty{
		AutomaticMaxWidth:  defaultAutomaticMaxWidth,
		Newline:            defaultNewline,
		Paragraph:          defaultParagraph,
		UseTrailingNewLine: defaultUseTrailingNewLine,

		Borders: &Borders{
			AsymmetricalDouble: BorderType{
				Horizontal:  "─",
				Vertical:    "║",
				TopLeft:     "╓",
				TopRight:    "╖",
				BottomLeft:  "╙",
				BottomRight: "╜",
			},
			AsymmetricalThick: BorderType{
				Horizontal:  "─",
				Vertical:    "┃",
				TopLeft:     "┎",
				TopRight:    "┒",
				BottomLeft:  "┖",
				BottomRight: "┚",
			},
			Color: defaultBorderColor,
			Double: BorderType{
				Horizontal:  "═",
				Vertical:    "║",
				TopLeft:     "╔",
				TopRight:    "╗",
				BottomLeft:  "╚",
				BottomRight: "╝",
			},
			Rounded: BorderType{
				Horizontal:  "─",
				Vertical:    "│",
				TopLeft:     "╭",
				TopRight:    "╮",
				BottomLeft:  "╰",
				BottomRight: "╯",
			},
			Square: BorderType{
				Horizontal:  "─",
				Vertical:    "│",
				TopLeft:     "┌",
				TopRight:    "┐",
				BottomLeft:  "└",
				BottomRight: "┘",
			},
			Thick: BorderType{
				Horizontal:  "━",
				Vertical:    "┃",
				TopLeft:     "┏",
				TopRight:    "┓",
				BottomLeft:  "┗",
				BottomRight: "┛",
			},
			Title: &Titles{
				Bold:          defaultBorderTitleBold,
				Case:          defaultBorderTitleCase,
				Color:         defaultBorderTitleColor,
				Italics:       defaultBorderTitleItalics,
				Spaced:        defaultBorderTitleSpaced,
				Underline:     defaultBorderTitleUnderline,
				UseTitleColor: defaultBorderUseTitleColor,
			},
		},

		OutputMargin: &OutputMargin{
			Character: defaultMarginChar,
			Width:     defaultMarginWidth,
		},

		OutputPadding: &OutputPadding{
			Character: defaultPaddingChar,
			Width:     defaultPaddingWidth,
		},

		Terminal: &Terminal{
			Width:  -1,
			Height: -1,
		},

		Titles: &Titles{
			Bold:          defaultTitleBold,
			Case:          defaultTitleCase,
			Color:         defaultTitleColor,
			Italics:       defaultTitleItalics,
			Spaced:        defaultTitleSpaced,
			Underline:     defaultTitleUnderline,
			UseTitleColor: defaultUseTitleColor,
		},

		WordWrap: &WordWrap{
			Breakpoints: defaultBreakpoints,
			BreakWord:   defaultBreakWord,
			Hyphen:      defaultHyphen,
		},
	}
}

// Wrap - Shorthand for declaring a new default instance of Pretty and calling its Wrap method
// which wraps one or more lines of text at the given length limit.
//
// limit: An int declaring how many characters long the max-width of the output will be.
// s:     One or more sets of strings to be wrapped. Separated by comma.
func Wrap(limit int, s ...string) []string {
	return New().Wrap(limit, s)
}

// Padding - Shorthand for declaring a new default instance of Pretty and calling its Padding method
// which adds symmetrical padding around the strings.
//
// s:  An array of strings to be padded.
func Padding(s ...string) []string {
	return New().Padding(s...)
}

// Center - Shorthand for declaring a new default instance of Pretty and calling its Center method
// which adds padding around the strings to make them seem center aligned.
//
// s:  An array of strings to be center aligned.
func Center(s ...string) []string {
	return New().Center(s...)
}

// BorderSquare - Shorthand for declaring a new default instance of Pretty and calling its Border method with the Square BorterType parameter
// which adds a regular square border around your text.
//
// s:  An array of strings to be ...
func BorderSquare(t string, s ...string) string {
	p := New()
	return p.Border(p.Borders.Square, t, s...)
}

// CombineStrings - Shorthand for declaring a new default instance of Pretty and calling its CombineStrings method
// which combines an array of strings, treating each string as a separate paragraph.
//
// strs: An array of strings to be combined.
func CombineStrings(strs ...string) string {
	return New().CombineStrings(strs...)
}
