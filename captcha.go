package captcha

import "fmt"

type Captcha interface {
	CreateByText(*OptionText) Result
	CreateByMath(*OptionMath) Result
}

type Result struct {
	Text string
	Data string
}

func CreateByText(option OptionText) (*Result, error) {
	var text string
	opt := getOptionByText(option)
	if opt.Text != nil {
		text = *opt.Text
	} else {
		text = opt.randomText()
	}

	data, err := createCaptcha(text, opt)
	if err != nil {
		return nil, err
	}

	return &Result{
		Text: text,
		Data: data,
	}, err
}

func createCaptcha(text string, opt *option) (string, error) {
	width := widthDefault
	if opt.Width != nil {
		width = uint8(*opt.Width)
	}

	height := heightDefault
	bg := ""
	if opt.Height != nil {
		height = uint8(*opt.Height)
	}
	if opt.BackgroundColor != nil {
		opt.IsColor = setBooleanTrue()
		bg = *opt.BackgroundColor
	}

	result := fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"%v\" height=\"%v\" viewBox=\"0,0,%v,%v\" style=\"transform: rotateX(180deg)\">", width, height, width, height)
	if bg != "" {
		result += fmt.Sprintf("<rect fill=\"%v\" width=\"100%%\" height=\"100%%\"/>", bg)
	}

	lineNoise := opt.drawLineNoise()
	pathText, err := opt.drawText(text)
	if err != nil {
		return "", err
	}
	result += fmt.Sprintf("%v%v</svg>", lineNoise, pathText)

	return result, nil
}

func setBooleanTrue() *bool {
	result := true
	return &result
}
