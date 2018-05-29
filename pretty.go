package pretty

const githubURL = ""
const bugReportURL = ""

const (
	defaultBreakpoints        = " -"
	defaultNewline            = "\n"
	defaultParagraph          = "\n\n"
	defaultBreakWord          = true
	defaultHyphen             = "-"
	defaultUseTrailingNewLine = false
	defaultPaddingChar        = "\u00A0"
	defaultPaddingWidth       = 3
	defaultMarginChar         = "\u00A0"
	defaultMarginWidth        = 2
)

// New returns a new instance of Pretty, initialized with the default settings.
func New() Pretty {
	return Pretty{
		Newline:            defaultNewline,
		Paragraph:          defaultParagraph,
		UseTrailingNewLine: defaultUseTrailingNewLine,
		Terminal: &Terminal{
			Width:  -1,
			Height: -1,
		},
		WordWrap: &WordWrap{
			Breakpoints: defaultBreakpoints,
			BreakWord:   defaultBreakWord,
			Hyphen:      defaultHyphen,
		},
		OutputPadding: &OutputPadding{
			Width:     defaultPaddingWidth,
			Character: defaultPaddingChar,
		},
		OutputMargin: &OutputMargin{
			Width:     defaultMarginWidth,
			Character: defaultMarginChar,
		},
		Borders: &Borders{
			Square: &BorderType{
				Horizontal:  "─",
				Vertical:    "│",
				TopLeft:     "┌",
				TopRight:    "┐",
				BottomLeft:  "└",
				BottomRight: "┘",
			},
			Rounded: &BorderType{
				Horizontal:  "─",
				Vertical:    "│",
				TopLeft:     "╭",
				TopRight:    "╮",
				BottomLeft:  "╰",
				BottomRight: "╯",
			},
			Double: &BorderType{
				Horizontal:  "═",
				Vertical:    "║",
				TopLeft:     "╔",
				TopRight:    "╗",
				BottomLeft:  "╚",
				BottomRight: "╝",
			},
			AsymmetricalDouble: &BorderType{
				Horizontal:  "─",
				Vertical:    "║",
				TopLeft:     "╓",
				TopRight:    "╖",
				BottomLeft:  "╙",
				BottomRight: "╜",
			},
			Thick: &BorderType{
				Horizontal:  "━",
				Vertical:    "┃",
				TopLeft:     "┏",
				TopRight:    "┓",
				BottomLeft:  "┗",
				BottomRight: "┛",
			},
			AsymmetricalThick: &BorderType{
				Horizontal:  "─",
				Vertical:    "┃",
				TopLeft:     "┎",
				TopRight:    "┒",
				BottomLeft:  "┖",
				BottomRight: "┚",
			},
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

// CombineStrings - Shorthand for declaring a new default instance of Pretty and calling its CombineStrings method
// which combines an array of strings, treating each string as a separate paragraph.
//
// strs: An array of strings to be combined.
func CombineStrings(strs ...string) string {
	return New().CombineStrings(strs...)
}
