package captcha

type OptionText struct {
	Size             *uint8
	Width            *uint16
	Height           *uint16
	FontSize         *uint8
	IsColor          *bool
	IsInverse        *bool
	Noise            *uint8
	BackgroundColor  *string
	IgnoreCharacters *string
	Text             *string
	CharactersPreset *string
}

type OptionMath struct {
	Size            *uint8
	Width           *uint16
	Height          *uint16
	FontSize        *uint8
	IsColor         *bool
	IsInverse       *bool
	Noise           *uint8
	BackgroundColor *string
	MathOperator    *matchOperator
	MathMin         *uint8
	MathMax         *uint16
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
	Text             *string
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
		Text:             opt.Text,
		IgnoreCharacters: opt.IgnoreCharacters,
		CharactersPreset: opt.CharactersPreset,
	}
}

func getOptionByMath(opt OptionMath) *option {
	return &option{
		Size:            opt.Size,
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
