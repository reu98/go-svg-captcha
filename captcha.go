// Package captcha provides an easy to use
package captcha

import (
	"fmt"
	"image/color"
)

type Result struct {
	// A random string or the result of an operation.
	Text string

	// The HTML code snippet for SVG.
	Data string
}

// CreateByText: generate a new captcha
func CreateByText(option OptionText) (*Result, error) {
	opt := getOptionByText(option)
	text := opt.text
	if text == "" {
		text = opt.randomText()
	}

	data, err := opt.createCaptcha(text)
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

	resultMath := opt.generateMathOperation()
	data, err := opt.createCaptcha(resultMath.Equation)
	if err != nil {
		return nil, err
	}

	return &Result{
		Data: data,
		Text: (*resultMath).Text,
	}, nil
}

func (opt *option) createCaptcha(text string) (string, error) {
	result := fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"%v\" height=\"%v\" viewBox=\"0,0,%v,%v\" style=\"transform: rotateX(180deg)\">", opt.width, opt.height, opt.width, opt.height)
	if opt.backgroundColor != color.Transparent {
		r, g, b, a := opt.backgroundColor.RGBA()
		result += fmt.Sprintf("<rect fill=\"rgba(%v, %v, %v, %v)\" width=\"100%%\" height=\"100%%\"/>", r>>8, g>>8, b>>8, a>>8)
	}

	lineNoise := opt.drawCurve()
	noise := opt.drawNoise()
	pathText, err := opt.drawText(text)
	if err != nil {
		return "", err
	}
	result += fmt.Sprintf("%v%v%v</svg>", lineNoise, pathText, noise)

	return result, nil
}
