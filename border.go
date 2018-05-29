package pretty

import (
	"strings"
	"unicode/utf8"
)

// TODO: Finish this function
func (p Pretty) Border(b BorderType, s ...string) string {
	var ret string
	var border []string
	var width int

	margin := repeatChar(p.OutputMargin.Character, p.OutputMargin.Width)

	newStr := p.CombineStrings(s...)

	width = getLongestStringLength(strings.SplitAfter(newStr, "\n"))

	top := margin + b.TopLeft + repeatChar(b.Horizontal, width+p.OutputPadding.Width-utf8.RuneCountInString(b.TopLeft)-utf8.RuneCountInString(b.TopRight)) + b.TopRight + margin + "\n\n"
	bottom := "\n\n" + margin + b.BottomLeft + repeatChar(b.Horizontal, width+p.OutputPadding.Width-utf8.RuneCountInString(b.BottomLeft)-utf8.RuneCountInString(b.BottomRight)) + b.BottomRight + margin
	border = append(border, top)

	border = append(border, bottom)

	return ret
}

func (p Pretty) addHorizontalBorder(s string, char string, w int) string {
	padding := repeatChar(p.OutputPadding.Character, p.OutputPadding.Width)
	if utf8.RuneCountInString(s) < w {

	}
	return char + padding + s + padding + char
}
