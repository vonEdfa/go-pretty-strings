package pretty

import (
	"strings"
	"unicode/utf8"
)

// WordWrap - Contains all the settings for word wrapping.
type WordWrap struct {
	// Breakpoints - Defines which characters should be able to break a line.
	// Default: " -"
	Breakpoints string

	// BreakWord - Defines if the wordwrapper is allowed to break unbreakable words that are too long when wrapping.
	// Default: true
	BreakWord bool

	// Hyphen - Defines which character(s) should be used to indicate that a word har been broken.
	// Default: "-"
	Hyphen string

	// UseTrailingNewLine - Defines if the output should end with a newline.
	// Default: false
	UseTrailingNewLine bool
}

// Wrap - Shorthand for declaring a new default instance of Pretty and calling its Wrap method.
//
// limit: An int declaring how many characters long the max-width of the output will be.
// s:     One or more sets of strings to be wrapped. The different strings will be treated as separate paragraphs. Separated by comma.
func Wrap(limit int, s ...string) []string {
	return New().Wrap(limit, s)
}

// Wrap - Wraps one or more lines of text at the given length limit.
//
// p:     Main Pretty struct containing all your settings.
// limit: An int declaring how many characters long the max-width of the output will be.
// strs:  An array of strings to be wrapped. The different strings will be treated as separate paragraphs.
func (p Pretty) Wrap(limit int, strs []string) []string {
	var ret []string

	// Loop through each string induvidually
	for _, str := range strs {
		var newStr string

		// // Start a new paraphraph if we have more than one string
		// if i > 0 {
		// 	ret = strings.TrimSuffix(ret, p.Newline)
		// 	ret += p.Paragraph
		// }

		// Look for newline characters and let's start by splitting there
		for _, s := range strings.Split(str, p.Newline) {
			// Clean the string from any leading or trailing spaces
			s = strings.Trim(s, " ")
			newStr += p.wrapLine(s, limit)

			if p.WordWrap.UseTrailingNewLine {
				newStr += p.Newline
			}
		}
		ret = append(ret, newStr)
	}

	return ret
}

func (p Pretty) wrapLine(s string, limit int) string {
	// We won't go lower than 1 character
	if limit < 1 {
		limit = 1
	}
	if utf8.RuneCountInString(s) < limit+1 {
		return s
	}

	i := strings.LastIndexAny(s[:limit+1], p.WordWrap.Breakpoints)
	// If no valid breakpoints were found...
	if i < 0 {
		// ...break word is allowed, break it!
		if p.WordWrap.BreakWord {
			hyp := p.WordWrap.Hyphen
			// Don't use the hyphen if we don't have space for it
			if i <= 1 {
				hyp = ""
			}
			i = limit - utf8.RuneCountInString(hyp)
			return s[:i] + hyp + p.Newline + p.wrapLine(s[i:], limit)
		}
		// ...otherwise wrap at the next breakpoint
		i = strings.IndexAny(s, p.WordWrap.Breakpoints)

		// When out of options, just complete.
		if i < 0 {
			return s
		}
	}

	return s[:i] + p.Newline + p.wrapLine(s[i+1:], limit)
}
