package pretty

import "strings"

// OutputPadding contains all the settings for how paddings should look.
type OutputPadding struct {
	// Width - Defines the width of the padding.
	// Default: 3
	Width int

	// Character - Defines the character used to fill the padding with.
	// Default: "\u00A0"
	Character string
}

func (p Pretty) Padding(strs []string) []string {
	var ret []string
	var right int

	left := p.OutputPadding.Width

	// Loop through each string induvidually
	for _, str := range strs {
		// Split by newline
		for _, s := range strings.Split(str, p.Newline) {
			p.addPaddingToLine(s, left, right)
		}
	}

	return ret
}

func (p Pretty) addPaddingToLine(s string, paddingLeft int, paddingRight int) string {
	left := repeatChar(p.OutputPadding.Character, paddingLeft)
	right := repeatChar(p.OutputPadding.Character, paddingRight)

	return left + s + right
}
