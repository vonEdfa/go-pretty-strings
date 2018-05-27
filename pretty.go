package pretty

import "strings"

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

var bordersRound = map[string]string{
	"topL": "╭",
	"topR": "╮",
	"hor":  "─",
	"ver":  "│",
	"botL": "╰",
	"botR": "╯",
}
var bordersSquare = map[string]string{
	"topL": "┌",
	"topR": "┐",
	"hor":  "─",
	"ver":  "│",
	"botL": "└",
	"botR": "┘",
}

// Pretty - Contains all the settings for customizable text formatting.
type Pretty struct {
	// Newline - Defines which characters should be used to create a new line.
	// Default: "\n"
	Newline string

	// Paragraph - Defines which characters should be used to define a new paragraph.
	// Default: "\n\n"
	Paragraph string

	// WordWrap - Contains all the settings for word wrapping.
	WordWrap *WordWrap

	// OutputPadding - Contains all the settings for how paddings should look.
	OutputPadding *OutputPadding

	// OutputMargin - Contains all the settings for how margins should look.
	OutputMargin *OutputMargin
}

// OutputMargin contains all the settings for how margin should look.
type OutputMargin struct {
	width     int
	character string
}

// New returns a new instance of Pretty, initialized with the default settings.
func New() Pretty {
	return Pretty{
		Newline:   defaultNewline,
		Paragraph: defaultParagraph,
		WordWrap: &WordWrap{
			Breakpoints:        defaultBreakpoints,
			BreakWord:          defaultBreakWord,
			Hyphen:             defaultHyphen,
			UseTrailingNewLine: defaultUseTrailingNewLine,
		},
		OutputPadding: &OutputPadding{
			Width:     defaultPaddingWidth,
			Character: defaultPaddingChar,
		},
		OutputMargin: &OutputMargin{
			width:     defaultMarginWidth,
			character: defaultMarginChar,
		},
	}
}

// CombineStrings - Shorthand for declaring a new default instance of Pretty and calling its CombineStrings method.
//
// strs: An array of strings to be combined.
func CombineStrings(strs []string) string {
	return New().CombineStrings(strs)
}

// CombineStrings - Combines an array of strings, treating each string as a separate paragraph.
//
// p:    Main Pretty struct containing all your settings.
// strs: An array of strings to be combined.
func (p Pretty) CombineStrings(strs []string) string {
	var ret string

	for i, s := range strs {
		if i > 0 {
			ret += p.Paragraph
		}
		ret += s
	}

	if p.WordWrap.UseTrailingNewLine && !strings.HasSuffix(ret, p.Newline) {
		ret += p.Newline
	}

	return ret
}

func getLongestStringLength(list []string) int {
	var longest int

	for _, s := range list {
		if len(s) > longest {
			longest = len(s)
		}
	}

	return longest
}

func repeatChar(char string, amount int) string {
	var ret string

	for i := 0; i < amount; i++ {
		ret += char
	}

	return ret
}
