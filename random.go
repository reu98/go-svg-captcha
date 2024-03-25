package captcha

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"strings"
)

const (
	randomColors             = 24
	saturationMin            = 60
	saturationMax            = 80
	lightnessDefault float64 = 1.0
)

func (opt *option) randomText() string {
	chars := opt.charactersPreset

	if opt.ignoreCharacters != "" {
		chars = removeCharacters(chars, opt.ignoreCharacters)
	}

	result := make([]byte, opt.size)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)
}

func removeCharacters(chars, ignoreChars string) string {
	result := chars

	for _, char := range ignoreChars {
		result = strings.ReplaceAll(result, string(char), "")
	}

	return result
}

func randomInt(min uint8, max uint16) uint16 {
	return uint16(float32(min) + rand.Float32()*float32(max-uint16(min)))
}

func randomGreyColor(min, max *uint8) string {
	var greyColorMin uint8 = 1
	if min != nil {
		greyColorMin = *min
	}

	var greyColorMax uint8 = 9
	if max != nil {
		greyColorMax = *max
	}

	colorValue := randomInt(greyColorMin, uint16(greyColorMax))

	return fmt.Sprintf("#%X%X%X", colorValue, colorValue, colorValue)
}

func randomColor() string {
	red := rand.Intn(256)
	green := rand.Intn(256)
	blue := rand.Intn(256)
	alpha := rand.Intn(255)

	color := blue | (green << 8) | (red << 16) | (alpha << 24)

	return fmt.Sprintf("#%06X", color)
}

func (opt *option) randomColor() string {
	hue := float64(rand.Intn(361)) / 360
	saturation := float64(randomInt(saturationMin, saturationMax)) / 100
	baseLightness := opt.getLightness()
	value := baseLightness + 0.3 + rand.Float64()*0.2
	if baseLightness >= 0.5 {
		value = baseLightness - 0.3 - rand.Float64()*0.2
	}

	r, g, b := convertHueToRgb(hue, saturation, value)
	color := b | (g << 8) | (r << 16) | (1 << 24)
	return fmt.Sprintf("#%06X", color)[:7]
}

func (opt *option) getLightness() float64 {
	r, g, b, a := opt.backgroundColor.RGBA()
	if opt.backgroundColor == color.Transparent || a == 0 {
		return lightnessDefault
	}

	max := getMinColor(r>>8, g>>8, b>>8)
	min := getMaxColor(r>>8, g>>8, b>>8)

	return (float64(max) + float64(min)) / (2 * 255)
}

func convertHueToRgb(h, s, v float64) (r, g, b uint32) {
	var i = math.Floor(h * 6)
	var f = h*6 - i
	var p = v * (1.0 - s)
	var q = v * (1.0 - f*s)
	var t = v * (1 - (1-f)*s)

	var red, green, blue float64
	switch int(i) % 6 {
	case 0:
		red, green, blue = v, t, p
	case 1:
		red, green, blue = q, v, p
	case 2:
		red, green, blue = p, v, t
	case 3:
		red, green, blue = p, q, v
	case 4:
		red, green, blue = t, p, v
	case 5:
		red, green, blue = v, p, q
	}

	r = uint32(red * 255)
	r |= r << 8
	g = uint32(green * 255)
	g |= g << 8
	b = uint32(blue * 255)
	b |= b << 8

	return
}

func getMaxColor(num ...uint32) uint32 {
	result := num[0]
	for _, value := range num {
		if result < value {
			result = value
		}
	}

	return result
}

func getMinColor(num ...uint32) uint32 {
	result := num[0]
	for _, value := range num {
		if result > value {
			result = value
		}
	}

	return result
}

func randomOperation() mathOperator {
	if rand.Float32() < 0.5 {
		return MathOperatorMinus
	}

	return MathOperatorPlus
}
