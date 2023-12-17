package captcha

import (
	"fmt"
	"math/rand"

	"github.com/tdewolff/canvas"
)

const fileNameFont = "./fonts/Comismsh.ttf"

func (opt *option) drawText(text string) (string, error) {
	width := widthDefault
	if opt.Width != nil {
		width = uint8(*opt.Width)
	}

	height := heightDefault
	if opt.Height != nil {
		height = uint8(*opt.Height)
	}

	isColor := false
	if opt.IsColor != nil && *opt.IsColor {
		isColor = true
	}

	fillColorMin := fillColorMinDefault
	fillColorMax := fillColorMaxDefault
	if opt.IsInverse != nil && *opt.IsInverse {
		fillColorMin = fillColorMinInverse
		fillColorMax = fillColorMaxInverse
	}

	letterWidth := float32((width - paddingHorizontalDefault*2)) / float32(len(text))
	result := ""
	for index, char := range text {
		var fill string
		if isColor {
			fill = randomColor(opt.BackgroundColor)
		} else {
			fill = randomGreyColor(uint8(fillColorMin), uint8(fillColorMax))
		}

		x := letterWidth*float32(index) + letterWidth/2
		y := float32(height / 2)
		d, err := opt.drawChar(char, x, y)
		if err != nil {
			return "", err
		}
		result += fmt.Sprintf("<path fill=\"%v\" d=\"%v\"/>", fill, d)
	}

	return result, nil
}

func (opt *option) drawChar(char rune, x float32, y float32) (string, error) {
	fontFamily, err := canvas.LoadFontFile(fileNameFont, canvas.FontRegular)
	if err != nil {
		return "", err
	}

	fontSize := fontSizeDefault
	if opt.FontSize != nil && *opt.FontSize > 0 {
		fontSize = *opt.FontSize
	}

	face := fontFamily.Face(float64(fontSize), nil)

	path, _, err := face.ToPath(string(char))
	if err != nil {
		return "", err
	}

	return opt.randomTranslatePath(path, x, y), nil
}

func (opt *option) drawLineNoise() string {
	index := 0
	noise := noiseDefault
	if opt.Noise != nil && *opt.Noise > noiseDefault {
		noise = *opt.Noise
	}

	min := noiseGreyColorMinDefault
	max := noiseGreyColorMaxDefault
	if opt.IsInverse != nil && *opt.IsInverse {
		min = noiseGreyColorMinInverse
		max = noiseGreyColorMaxInverse
	}

	width := widthDefault
	if opt.Width != nil {
		width = uint8(*opt.Width)
	}

	height := heightDefault
	if opt.Height != nil {
		height = uint8(*opt.Height)
	}

	result := ""
	for uint8(index) < noise {
		var stroke string
		if opt.IsColor != nil && *opt.IsColor {
			stroke = randomColor(opt.BackgroundColor)
		} else {
			stroke = randomGreyColor(min, max)
		}

		moveLine := fmt.Sprintf("%v %v", randomInt(1, 21), randomInt(1, uint16(height)-1))
		cubicStart := fmt.Sprintf("%v %v", randomInt(width/2-21, uint16(width/2+21)), randomInt(1, uint16(height-1)))
		cubicMid := fmt.Sprintf("%v %v", randomInt(width/2-21, uint16(width/2+21)), randomInt(1, uint16(height-1)))
		cubicEnd := fmt.Sprintf("%v %v", randomInt(width-21, uint16(width-1)), randomInt(1, uint16(height-1)))
		result += fmt.Sprintf("<path d=\"M%v C%v,%v,%v\" stroke=\"%v\" fill=\"none\"/> ", moveLine, cubicStart, cubicMid, cubicEnd, stroke)
		index++
	}

	return result
}

func (opt *option) randomTranslatePath(p *canvas.Path, x, y float32) string {
	x = opt.calculateRandomOffset(x)
	y = opt.calculateRandomOffset(y)
	return p.Translate(float64(x), float64(y)).String()
}

func (opt *option) calculateRandomOffset(offset float32) float32 {
	operation := opt.randomOperation()

	if operation == MathOperatorPlus {
		return offset + opt.randomOffset()
	}

	return offset - opt.randomOffset()
}

func (opt *option) randomOperation() matchOperator {
	if rand.Float32() < 0.5 {
		return MathOperatorMinus
	}

	return MathOperatorPlus
}

func (opt *option) randomOffset() float32 {
	return (rand.Float32() * 0.2) - 0.1
}
