package pretty

import (
	"math"
	"strings"
	"unicode/utf8"
)

// Padding - Adds symmetrical padding around the strings.
//
// p:     Main Pretty struct containing all your settings.
// strs:  An array of strings to be padded.
func (p Pretty) Padding(strs ...string) []string {
	var ret []string
	var right int

	left := p.OutputPadding.Width

	// Loop through each string induvidually
	for _, str := range strs {
		var newStr string
		// Split by newline
		for i, s := range strings.Split(str, p.Newline) {
			newStr += p.addPaddingToLine(s, left, right)
			if i < len(strings.Split(str, p.Newline))-1 || p.UseTrailingNewLine {
				newStr += "\n"
			}
		}
		ret = append(ret, newStr)
	}

	return ret
}

// Center - Adds padding around the strings to make them seem center aligned.
//
// p:     Main Pretty struct containing all your settings.
// strs:  An array of strings to be center aligned.
func (p Pretty) Center(strs ...string) []string {
	var ret []string

	for _, str := range strs {
		var newStr string
		var left int
		var right int
		width := getLongestStringLength(strings.Split(str, p.Newline))

		for i, s := range strings.Split(str, p.Newline) {
			if utf8.RuneCountInString(s) < width {
				diff := width - utf8.RuneCountInString(s)
				if diff%2 == 0 {
					left = diff / 2
					right = left
				} else {
					leftf := math.Floor(float64(diff) / 2)
					rightf := math.Ceil(float64(diff) / 2)
					left = int(leftf)
					right = int(rightf)
				}
				newStr += p.addPaddingToLine(s, left, right)
			} else {
				newStr += s
			}

			if i < len(strings.Split(str, p.Newline))-1 || p.UseTrailingNewLine {
				newStr += "\n"
			}
		}
		ret = append(ret, newStr)
	}

	return ret
}

func (p Pretty) addPaddingToLine(s string, paddingLeft int, paddingRight int) string {
	left := repeatChar(p.OutputPadding.Character, paddingLeft)
	right := repeatChar(p.OutputPadding.Character, paddingRight)

	return left + s + right
}
