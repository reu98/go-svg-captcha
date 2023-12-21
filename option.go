package captcha

type OptionText struct {
	// The length of the random string
	// Default: 4
	Size *uint8

	// The width of the SVG captcha
	// Default: 150
	Width *uint16

	// The height of the SVG captcha
	// Default: 50
	Height *uint16

	// The font size in the captcha
	// Default: 12
	FontSize *uint8

	// The color of the characters in the captcha.
	// True: the characters will have individual colors.
	// False: The characters will be gray
	// Default: false
	IsColor *bool

	// Invert the colors.
	// Default: false
	IsInverse *bool

	// The number of lines in the captcha
	// Default: 1
	Noise *uint8

	// Background color of the SVG captcha
	BackgroundColor *string

	// Remove unacceptable characters from the captcha.
	IgnoreCharacters *string

	// The characters that can be displayed in the captcha.
	// Default: ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	CharactersPreset *string
}

type OptionMath struct {
	// The width of the SVG captcha
	// Default: 150
	Width *uint16

	// The height of the SVG captcha
	// Default: 50
	Height *uint16

	// The font size in the captcha
	// Default: 12
	FontSize *uint8

	// The color of the characters in the captcha.
	// True: the characters will have individual colors.
	// False: The characters will be gray
	// Default: false
	IsColor *bool

	// Invert the colors.
	// Default: false
	IsInverse *bool

	// The number of lines in the captcha
	// Default: 1
	Noise *uint8

	// Background color of the SVG captcha
	BackgroundColor *string

	// The operation for the mathematical calculation
	// Supports addition (+) and subtraction (-) operations
	// If there's no specified operation, it will automatically choose one.
	MathOperator *matchOperator

	// The minimum value for a number in the operation
	// Default: 1
	MathMin *uint8

	// The maximum value for a number in the operation.
	// Default: 9
	MathMax *uint16
}

type option struct {
	Size             *uint8
	Width            *uint16
	Height           *uint16
	FontSize         *uint8
	IsColor          *bool
	IsInverse        *bool
	Noise            *uint8
	BackgroundColor  *string
	MathOperator     *matchOperator
	MathMin          *uint8
	MathMax          *uint16
	IgnoreCharacters *string
	CharactersPreset *string
}

func getOptionByText(opt OptionText) *option {
	return &option{
		Size:             opt.Size,
		Width:            opt.Width,
		Height:           opt.Height,
		FontSize:         opt.FontSize,
		IsColor:          opt.IsColor,
		IsInverse:        opt.IsInverse,
		Noise:            opt.Noise,
		BackgroundColor:  opt.BackgroundColor,
		IgnoreCharacters: opt.IgnoreCharacters,
		CharactersPreset: opt.CharactersPreset,
	}
}

func getOptionByMath(opt OptionMath) *option {
	return &option{
		Width:           opt.Width,
		Height:          opt.Height,
		FontSize:        opt.FontSize,
		IsColor:         opt.IsColor,
		IsInverse:       opt.IsInverse,
		Noise:           opt.Noise,
		BackgroundColor: opt.BackgroundColor,
		MathOperator:    opt.MathOperator,
		MathMin:         opt.MathMin,
		MathMax:         opt.MathMax,
	}
}
