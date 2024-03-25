package captcha

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/reu98/go-svg-captcha/fonts"
	"github.com/tdewolff/canvas"
)

func (opt *option) drawText(text string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()

	fillColorMin := fillColorMinDefault
	fillColorMax := fillColorMaxDefault
	if opt.isInverse {
		fillColorMin = fillColorMinInverse
		fillColorMax = fillColorMaxInverse
	}

	letterWidth := float32((opt.width - uint16(paddingHorizontalDefault)*2)) / float32(len(text))
	result := ""
	for index, char := range text {
		var fill string
		if opt.isColor {
			fill = opt.randomColor()
		} else {
			min := uint8(fillColorMin)
			max := uint8(fillColorMax)
			fill = randomGreyColor(&min, &max)
		}

		minY := float32(opt.height) / 4
		maxY := float32(opt.height) / 2
		x := letterWidth*float32(index) + letterWidth/2
		y := minY + rand.Float32()*(maxY-minY)
		d, err := opt.drawChar(char, x, y)
		if err != nil {
			return "", err
		}
		result += fmt.Sprintf("<path fill=\"%v\" d=\"%v\"/>", fill, d)
	}

	return result, nil
}

func (opt *option) drawNoise() string {
	var result string
	totalNoise := (opt.width * opt.height) / 28
	for i := 0; i < int(totalNoise); i++ {
		x := rand.Intn(int(opt.width))
		y := rand.Intn(int(opt.height))
		color := randomColor()

		result += fmt.Sprintf("<circle cx=\"%d\" cy=\"%d\" r=\"%d\" fill=\"%s\" />", x, y, noiseRadius, color)
	}

	return result
}

func (opt *option) drawChar(char rune, x float32, y float32) (string, error) {
	fontFamily, err := loadFont(opt.fontPath)
	if err != nil {
		return "", err
	}

	fontFace := fontFamily.Face(float64(opt.fontSize), nil)
	path, _, err := fontFace.ToPath(string(char))
	if err != nil {
		return "", err
	}

	return randomTranslatePath(path, x, y), nil
}

func loadFont(fontPath string) (*canvas.Font, error) {
	if fontPath != "" {
		fontFamily, err := canvas.LoadFontFile(fontPath, canvas.FontRegular)
		if err != nil {
			return nil, err
		}

		return fontFamily, nil
	}

	fontFamily, err := canvas.LoadFont(fonts.Comismsh, 0, canvas.FontRegular)
	if err != nil {
		return nil, err
	}

	return fontFamily, nil
}

func (opt *option) drawCurve() string {
	count := opt.curve
	var result string
	for i := 0; i < int(count); i++ {
		stroke := opt.randomColor()
		moveLine := fmt.Sprintf("%v %v", randomInt(1, 21), randomInt(1, uint16(opt.height)-1))
		cubicStart := fmt.Sprintf("%v %v", randomInt(uint8(opt.width/2-21), uint16(opt.width/2+21)), randomInt(1, uint16(opt.height-1)))
		cubicMid := fmt.Sprintf("%v %v", randomInt(uint8(opt.width/2-21), uint16(opt.width/2+21)), randomInt(1, uint16(opt.height-1)))
		cubicEnd := fmt.Sprintf("%v %v", randomInt(uint8(opt.width-21), uint16(opt.width-1)), randomInt(1, uint16(opt.height-1)))

		result += fmt.Sprintf("<path d=\"M%v C%v,%v,%v\" stroke=\"%v\" fill=\"none\"/> ", moveLine, cubicStart, cubicMid, cubicEnd, stroke)
	}

	return result
}

func randomTranslatePath(p *canvas.Path, x, y float32) string {
	x = calculateRandomOffset(x)
	y = calculateRandomOffset(y)
	return p.Translate(float64(x), float64(y)).String()
}

func calculateRandomOffset(offset float32) float32 {
	operation := randomOperation()

	if operation == MathOperatorPlus {
		return offset + randomOffset()
	}

	return offset - randomOffset()
}

func randomOffset() float32 {
	return (rand.Float32() * 0.2) - 0.1
}
