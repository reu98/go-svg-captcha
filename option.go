package captcha

import (
	"image/color"
)

type OptionText struct {
	// The length of the random string
	// Default: 4
	Size uint8

	// The width of the SVG captcha
	// Default: 150
	Width uint16

	// The height of the SVG captcha
	// Default: 50
	Height uint16

	// The font size in the captcha
	// Default: 16
	FontSize uint8

	// The color of the characters in the captcha.
	// True: the characters will have individual colors.
	// False: The characters will be gray
	// Default: false
	IsColor bool

	// Invert the colors.
	// Default: false
	IsInverse bool

	// The number of lines in the captcha
	// Default: 1
	Curve uint8

	// Background color of the SVG captcha
	// Default: Transperant
	BackgroundColor color.Color

	// Remove unacceptable characters from the captcha.
	IgnoreCharacters string

	// Path font
	FontPath string

	// The characters that can be displayed in the captcha.
	// Default: ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	CharactersPreset string

	// Generate With Text
	// Note: Text is optional. If not provided, it will be generated randomly
	Text string
}

type OptionMath struct {
	// The width of the SVG captcha
	// Default: 150
	Width uint16

	// The height of the SVG captcha
	// Default: 50
	Height uint16

	// The font size in the captcha
	// Default: 16
	FontSize uint8

	// Path font
	FontPath string

	// The color of the characters in the captcha.
	// True: the characters will have individual colors.
	// False: The characters will be gray
	// Default: false
	IsColor bool

	// Invert the colors.
	// Default: false
	IsInverse bool

	// The number of lines in the captcha
	// Default: 1
	Curve uint8

	// Background color of the SVG captcha
	// Default: Transperant
	BackgroundColor color.Color

	// The operation for the mathematical calculation
	// Supports addition (+) and subtraction (-) operations
	// If there's no specified operation, it will automatically choose one.
	MathOperator mathOperator

	// The minimum value for a number in the operation
	// Default: 1
	MathMin uint8

	// The maximum value for a number in the operation.
	// Default: 9
	MathMax uint16
}

type option struct {
	size             uint8
	text             string
	width            uint16
	height           uint16
	isColor          bool
	isInverse        bool
	curve            uint8
	backgroundColor  color.Color
	mathOperator     mathOperator
	mathMin          uint8
	mathMax          uint16
	ignoreCharacters string
	charactersPreset string
	fontSize         uint
	fontPath         string
}

func getOptionByText(opt OptionText) *option {
	size := sizeDefault
	if opt.Size != 0 {
		size = opt.Size
	}

	width := widthDefault
	if opt.Width != 0 {
		width = opt.Width
	}

	height := heightDefault
	if opt.Height != 0 {
		height = opt.Height
	}

	fontSize := fontSizeDefault
	if opt.FontSize != 0 {
		fontSize = opt.FontSize
	}

	curve := curveDefault
	if opt.Curve != 0 {
		curve = opt.Curve
	}

	isColor := opt.IsColor
	var backgroundColor color.Color = color.Transparent
	if opt.BackgroundColor != nil {
		backgroundColor = opt.BackgroundColor
		isColor = true
	}

	charPreset := characters
	if opt.CharactersPreset != "" {
		charPreset = opt.CharactersPreset
	}

	return &option{
		size:             size,
		width:            width,
		height:           height,
		fontSize:         uint(fontSize) * uint(ratioFontSize),
		isColor:          isColor,
		isInverse:        opt.IsInverse,
		curve:            curve,
		backgroundColor:  backgroundColor,
		ignoreCharacters: opt.IgnoreCharacters,
		charactersPreset: charPreset,
		fontPath:         opt.FontPath,
		text:             opt.Text,
	}
}

func getOptionByMath(opt OptionMath) *option {
	width := widthDefault
	if opt.Width != 0 {
		width = opt.Width
	}

	height := heightDefault
	if opt.Height != 0 {
		height = opt.Height
	}

	fontSize := fontSizeDefault
	if opt.FontSize != 0 {
		fontSize = opt.FontSize
	}

	curve := curveDefault
	if opt.Curve != 0 {
		curve = opt.Curve
	}

	isColor := opt.IsColor
	var backgroundColor color.Color = color.Transparent
	if opt.BackgroundColor != nil {
		backgroundColor = opt.BackgroundColor
		isColor = true
	}

	mathOperator := opt.MathOperator
	if mathOperator == "" {
		mathOperator = randomOperation()
	}

	mathMin := mathMinDefault
	if opt.MathMin != 0 {
		mathMin = opt.MathMin
	}

	mathMax := mathMaxDefault
	if opt.MathMax != 0 {
		mathMax = opt.MathMax
	}

	return &option{
		width:           width,
		height:          height,
		fontSize:        uint(fontSize) * uint(ratioFontSize),
		isColor:         isColor,
		curve:           curve,
		isInverse:       opt.IsInverse,
		backgroundColor: backgroundColor,
		mathOperator:    mathOperator,
		mathMin:         mathMin,
		mathMax:         mathMax,
		fontPath:        opt.FontPath,
	}
}
