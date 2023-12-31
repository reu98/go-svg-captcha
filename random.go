package captcha

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

const (
	randomColors             = 24
	saturationMin            = 60
	saturationMax            = 80
	lightnessDefault float32 = 1.0
)

func (opt *option) randomText() string {
	size := sizeDefault
	if opt.Size != nil {
		size = *opt.Size
	}

	chars := characters
	if opt.CharactersPreset != nil {
		chars = *opt.CharactersPreset
	}

	if opt.IgnoreCharacters != nil {
		chars = removeCharacters(chars, *opt.IgnoreCharacters)
	}

	result := make([]byte, size)
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

func randomColor(bgColor *string) string {
	if bgColor != nil {
		regexColor := regexp.MustCompile(regexColor)
		if !regexColor.MatchString(*bgColor) {
			bgColor = nil
		}
	}

	hue := float32(randomInt(0, uint16(randomColors))) / randomColors
	saturation := float32(randomInt(saturationMin, saturationMax)) / 100
	bgLightness := lightnessDefault
	if bgColor != nil {
		bgLightness = getLightness(*bgColor)
	}

	maxLightness := int(bgLightness*100) - 25
	minLightness := int(bgLightness*100) - 45
	if bgLightness < 0.5 {
		minLightness = int(bgLightness*100) + 25
		maxLightness = int(bgLightness*100) + 45
	}

	lightness := float32(randomInt(uint8(minLightness), uint16(maxLightness))) / 100
	calculateQ := lightness + saturation - (lightness * saturation)
	if lightness < 0.5 {
		calculateQ = lightness * (lightness + saturation)
	}

	calculateP := (2 * lightness) - calculateQ
	red := int(convertHueToRgb(hue+(1/3), calculateP, calculateQ) * 255)
	green := int(convertHueToRgb(hue, calculateP, calculateQ) * 255)
	blue := int(convertHueToRgb(hue-(1/3), calculateP, calculateQ) * 255)
	color := (blue | (green << 8) | (red << 16) | (1 << 24))
	hex := strconv.FormatInt(int64(color), 16)
	return fmt.Sprintf("#%v", hex[1:])
}

func getLightness(bgColor string) float32 {
	rgbColor := trimFirstRune(bgColor)
	if len(rgbColor) == 3 {
		rgbColor = fmt.Sprintf("%v%v%v%v%v%v", string(rgbColor[0]), string(rgbColor[0]), string(rgbColor[1]), string(rgbColor[1]), string(rgbColor[2]), string(rgbColor[2]))
	}

	hex, err := strconv.ParseInt(rgbColor, 16, 64)
	if err != nil {
		return lightnessDefault
	}

	red := hex >> 16
	green := (hex >> 8) & 255
	blue := hex & 255
	min := getMin(int(red), int(green), int(blue))
	max := getMax(int(red), int(green), int(blue))

	return float32((max + min) / (2 * 255))
}

func convertHueToRgb(hue, p, q float32) float32 {
	var one float32 = 1
	var two float32 = 2
	var three float32 = 3
	var six float32 = 6
	switch {
	case hue*six < one:
		return p + (q-p)*hue*six
	case hue*two < one:
		return q
	case hue*three < two:
		return p + (q-p)*((two/three)-hue)*six
	default:
		return p
	}
}

func trimFirstRune(s string) string {
	return strings.TrimPrefix(s, string(s[0]))
}

func getMax(num ...int) int {
	result := num[0]
	for _, value := range num {
		if result < value {
			result = value
		}
	}

	return result
}

func getMin(num ...int) int {
	result := num[0]
	for _, value := range num {
		if result > value {
			result = value
		}
	}

	return result
}
