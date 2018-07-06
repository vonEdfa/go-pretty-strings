package pretty

import (
	"strings"
	"unicode/utf8"
)

// FIXME: Rewrite and split function so we can esily support centered text too!
func (p Pretty) Border(b BorderType, t string, s ...string) string {
	var ret string
	var border []string
	var width int
	var top string
	var verPad string
	var bottom string

	margin := repeatChar(p.OutputMargin.Character, p.OutputMargin.Width)
	padding := repeatChar(p.OutputPadding.Character, p.OutputPadding.Width)

	newStr := p.CombineStrings(s...)

	width = getLongestStringLength(strings.SplitAfter(newStr, "\n"))

	if width < utf8.RuneCountInString(t) {
		width = utf8.RuneCountInString(t) + 1
	}

	if t == "" {
		top = margin + b.TopLeft + repeatChar(b.Horizontal, width+p.OutputPadding.Width*2) + b.TopRight + margin + "\n"
	} else {
		diff := width + p.OutputPadding.Width*2 - utf8.RuneCountInString(t)
		lineL := repeatChar(b.Horizontal, p.OutputPadding.Width-1)
		lineR := repeatChar(b.Horizontal, diff-p.OutputPadding.Width-1)
		title := p.addPaddingToLine(p.borderTitle(t), 1, 1)

		top = margin + b.TopLeft + lineL + title + lineR + b.TopRight + margin + "\n"
	}

	verPad = margin + b.Vertical + repeatChar(p.OutputPadding.Character, width+p.OutputPadding.Width*2) + b.Vertical + margin + "\n"
	bottom = margin + b.BottomLeft + repeatChar(b.Horizontal, width+p.OutputPadding.Width*2) + b.BottomRight + margin

	border = append(border, top, verPad)

	for _, line := range strings.Split(newStr, "\n") {
		var newL string

		if utf8.RuneCountInString(line) < width {
			// Add padding
			diff := width - utf8.RuneCountInString(line)
			fill := repeatChar(p.OutputPadding.Character, diff)
			newL = margin + b.Vertical + padding + line + fill + padding + b.Vertical + margin + "\n"
		} else if line == "\n" {
			border = append(border, verPad)
		} else {
			newL = margin + b.Vertical + padding + line + padding + b.Vertical + margin + "\n"
		}
		border = append(border, newL)
	}

	border = append(border, verPad, bottom)

	ret = strings.Join(border, "")

	return ret
}

func (p Pretty) borderTitle(t string) string {
	return p.FormatTitle(t, p.Borders.Title.Bold, p.Borders.Title.Italics, p.Borders.Title.Underline, p.Borders.Title.UseTitleColor, p.Borders.Title.Spaced)
}

func (p Pretty) addHorizontalBorder(s string, char string, w int) string {
	padding := repeatChar(p.OutputPadding.Character, p.OutputPadding.Width)
	if utf8.RuneCountInString(s) < w {

	}
	return char + padding + s + padding + char
}
