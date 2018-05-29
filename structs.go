package pretty

// Pretty - Contains all the settings for customizable text formatting.
type Pretty struct {
	// Newline - Defines which characters should be used to create a new line.
	// Default: "\n"
	Newline string

	// Paragraph - Defines which characters should be used to define a new paragraph.
	// Default: "\n\n"
	Paragraph string

	// UseTrailingNewLine - Defines if the output should end with a newline.
	// Default: false
	UseTrailingNewLine bool

	// Terminal - Contains all the info collected regarding the terminal running the program.
	Terminal *Terminal

	// WordWrap - Contains all the settings for word wrapping.
	WordWrap *WordWrap

	// OutputPadding - Contains all the settings for how paddings should look.
	OutputPadding *OutputPadding

	// OutputMargin - Contains all the settings for how margins should look.
	OutputMargin *OutputMargin

	// Border - Contains a set of predesigned border types when used together with the New function.
	Borders *Borders
}

// Terminal - Contains all the info collected regarding the terminal running the program.
type Terminal struct {
	// Width - The width of the current terminal window. (1 = one symbol)
	// Value should be set to -1 before calling the GetTermSize method.
	// Default: -1
	Width int

	// Height - The height of the current terminal window. (1 = one symbol)
	// Value should be set to -1 before calling the GetTermSize method.
	// Default: -1
	Height int
}

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
}

// OutputPadding contains all the settings for how paddings should look.
type OutputPadding struct {
	// Width - Defines the width of the padding.
	// Default: 3
	Width int

	// Character - Defines the character used to fill the padding with.
	// Default: "\u00A0"
	Character string
}

// OutputMargin - Contains all the settings for how margins should look.
type OutputMargin struct {
	Width     int
	Character string
}

// Borders - Contains a set of predesigned border types when used together with the New method.
type Borders struct {
	Square             *BorderType
	Rounded            *BorderType
	Double             *BorderType
	AsymmetricalDouble *BorderType
	Thick              *BorderType
	AsymmetricalThick  *BorderType
}

// BorderType - Struct to store all the symbols needed to draw a border.
type BorderType struct {
	// Horizontal - Defines the character used to create horizontal lines.
	Horizontal string

	// Vertical - Defines the character used to create vertical lines.
	Vertical string

	// TopLeft - Defines the character used to create the top left corner.
	TopLeft string

	// TopRight - Defines the character used to create the top right corner.
	TopRight string

	// BottomLeft - Defines the character used to create the bottom left corner.
	BottomLeft string

	// BottomRight - Defines the character used to create the bottom right corner.
	BottomRight string
}
