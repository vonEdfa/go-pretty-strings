package pretty

import (
	"strings"
)

func (p Pretty) Title(t string) string {
	return p.FormatTitle(t, p.Titles.Bold, p.Titles.Italics, p.Titles.Underline, p.Titles.UseTitleColor, p.Titles.Spaced)
}

func (p Pretty) FormatTitle(t string, isBold bool, isItalics bool, isUnderlined bool, useColor bool, isSpaced bool) string {
	var newTitle string
	var ansi string

	var escPre = "\033["
	var escSuf = "m"
	var b = "1;"
	var i = "3;"
	var u = "4;"
	var reset = "\033[0m"

	newTitle = t
	ansi = escPre

	if isBold {
		ansi += b
	} else {
		ansi += "0;"
	}

	if isItalics {
		ansi += i
	}

	if isUnderlined {
		ansi += u
	}

	if useColor {
		ansi += p.Titles.Color
	}

	ansi += escSuf

	newTitle = p.getCasedTitle(newTitle)

	if isSpaced {
		newTitle = getSpacedText(newTitle)
	}

	newTitle = ansi + newTitle + reset

	return newTitle
}

func (p Pretty) getCasedTitle(t string) string {
	switch p.Titles.Case {
	case UpperCase:
		return strings.ToUpper(t)
	case TitleCase:
		return strings.Title(t)
	case LowerCase:
		return strings.ToLower(t)
	default:
		return t
	}
}

func getSpacedText(t string) string {
	var ret string
	var newStr []string

	newStr = strings.Split(t, "")
	ret = strings.Join(newStr, " ")

	return ret
}
