// Package captcha provides an easy to use
package captcha

import "fmt"

type Result struct {
	// A random string or the result of an operation.
	Text string

	// The HTML code snippet for SVG.
	Data string
}

// CreateByText: generate a new captcha
func CreateByText(option OptionText) (*Result, error) {
	opt := getOptionByText(option)
	text := opt.randomText()

	data, err := createCaptcha(text, opt)
	if err != nil {
		return nil, err
	}

	return &Result{
		Text: text,
		Data: data,
	}, err
}

// CreateByMath: Generate a new captcha.
// It will return a captcha with an operation like 1 + 1.
func CreateByMath(option OptionMath) (*Result, error) {
	opt := getOptionByMath(option)
	min := mathMinDefault
	if opt.MathMin != nil {
		min = *opt.MathMin
	}

	max := mathMaxDefault
	if opt.MathMax != nil {
		max = *opt.MathMax
	}

	var operator matchOperator
	if opt.MathOperator != nil {
		operator = *opt.MathOperator
	} else {
		operator = randomOperation()
	}

	resultMath := generateMathOperation(&min, &max, &operator)
	data, err := createCaptcha((*resultMath).Equation, opt)
	if err != nil {
		return nil, err
	}

	return &Result{
		Data: data,
		Text: (*resultMath).Text,
	}, nil
}

func createCaptcha(text string, opt *option) (string, error) {
	width := widthDefault
	if opt.Width != nil {
		width = *opt.Width
	}

	height := heightDefault
	if opt.Height != nil {
		height = *opt.Height
	}

	bg := ""
	if opt.BackgroundColor != nil {
		isColor := true
		opt.IsColor = &isColor
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
