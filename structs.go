package pretty

/**
 * Keep the structs alphabetical please!
 * Thank you~
 */

// Pretty - Contains all the settings for customizable text formatting.
type Pretty struct {
	// AutomaticMaxWidth - Bool for if pretty should automatically declare the terminal with as the max width for all outputs.
	// Note: Only applicable if term width has been successfully calculated (i.e. is >= 0).
	// Default: true
	// TODO: Make sure this is implemented everywhere!
	// BUG: Test
	AutomaticMaxWidth bool

	// Newline - Defines which characters should be used to create a new line.
	// Default: "\n"
	Newline string

	// Paragraph - Defines which characters should be used to define a new paragraph.
	// Default: "\n\n"
	Paragraph string

	// UseTrailingNewLine - Defines if the output should end with a newline.
	// Default: false
	UseTrailingNewLine bool

	// Border - Contains a set of predesigned border types when used together with the New function.
	Borders *Borders

	// OutputMargin - Contains all the settings for how margins should look.
	OutputMargin *OutputMargin

	// OutputPadding - Contains all the settings for how paddings should look.
	OutputPadding *OutputPadding

	// Terminal - Contains all the info collected regarding the terminal running the program.
	Terminal *Terminal

	// Titles - Contains all the settings for how titles should look.
	Titles *Titles

	// WordWrap - Contains all the settings for word wrapping.
	WordWrap *WordWrap
}

// Borders - Contains a set of predesigned border types when used together with the New method.
type Borders struct {
	// AsymmetricalDouble - Square shaped with double lines on the sides, single on top.
	AsymmetricalDouble BorderType

	// AsymmetricalThick - Square shaped with thick lines on the sides, thin on top.
	AsymmetricalThick BorderType

	// Color - String containing an ANSI compatible color code. Leading escape characters or trailing letter "m" should not be included.
	// ANSI color code cheat sheet: https://gist.github.com/chrisopedia/8754917
	// Example: "30" for black.
	// Default: "37"
	// TODO: Pick a default color!
	Color string

	// Double - Square shaped with double lines.
	Double BorderType

	// Rounded - Regular border with rounded corners.
	Rounded BorderType

	// Square - Regular square border.
	Square BorderType

	// Thick - Square shaped with thick lines.
	Thick BorderType

	// Title - Contains all the settings for titles on bordered boxes.
	Title *Titles
}

// BorderType - Struct to store all the symbols needed to draw a border.
type BorderType struct {
	// BottomLeft - Defines the character used to create the bottom left corner.
	BottomLeft string

	// BottomRight - Defines the character used to create the bottom right corner.
	BottomRight string

	// Horizontal - Defines the character used to create horizontal lines.
	Horizontal string

	// TopLeft - Defines the character used to create the top left corner.
	TopLeft string

	// TopRight - Defines the character used to create the top right corner.
	TopRight string

	// Vertical - Defines the character used to create vertical lines.
	Vertical string
}

// OutputMargin - Contains all the settings for how margins should look.
type OutputMargin struct {
	// Character - Defines the character used to fill the padding with.
	// Default: "\u00A0"
	Character string

	// Width - Defines the width of the padding.
	// Default: 2
	Width int
}

// OutputPadding contains all the settings for how paddings should look.
type OutputPadding struct {
	// Character - Defines the character used to fill the padding with.
	// Default: "\u00A0"
	Character string

	// Width - Defines the width of the padding.
	// Default: 3
	Width int
}

// Terminal - Contains all the info collected regarding the terminal running the program.
type Terminal struct {
	// Height - The height of the current terminal window. (1 = one symbol)
	// Value should be set to -1 before calling the GetTermSize method.
	// Default: -1
	Height int

	// Width - The width of the current terminal window. (1 = one symbol)
	// Value should be set to -1 before calling the GetTermSize method.
	// Default: -1
	Width int
}

// Titles - Contains all the settings for how titles should look.
type Titles struct {
	// Bold - Bool for outputting the titles in bold (supported terminals only).
	// Default: true
	// TODO: Implement bold!
	Bold bool

	// Case - Int value deciding which case we want to use. Pretty uses enumerated public constants for your convenience.
	// 0 (InheritCase): Inherit (do nothing).
	// 1 (UpperCase): Uppercase.
	// 2 (TitleCase): Title case.
	// 3 (LowerCase): Lower case.
	// Default: TitleCase
	// TODO: Implement cases!
	Case int

	// Italics - Bool for outputting the titles in italics (supported terminals only).
	// Default: false
	// TODO: Implement italics!
	Italics bool

	// Spaced - Bool for if the title should be outputted in a spaced format:
	// Example: "M y   T i t l e"
	// Default: false
	// TODO: Implement spaced!
	Spaced bool

	// Underline - Bool for outputting the titles with an underline (supported terminals only).
	// Default: false
	// TODO: Implement underline!
	Underline bool

	// UseTitleColor - Bool for if titles should use the selected title color or not (supported terminals only).
	// If set to false it will use the defult terminal text color.
	// Default: true
	// TODO: Implement colors!
	UseTitleColor bool

	// Color - String containing an ANSI compatible color code. Leading escape characters or trailing letter "m" should not be included.
	// ANSI color code cheat sheet: https://gist.github.com/chrisopedia/8754917
	// Example: "30" for black.
	// Default: "37"
	// TODO: Pick a default color!
	Color string
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
