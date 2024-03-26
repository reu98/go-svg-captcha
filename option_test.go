package captcha

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetOptionByText(t *testing.T) {
	option := OptionText{}
	data := getOptionByText(option)

	require.NotNil(t, data)
	require.Equal(t, data.size, sizeDefault)
	require.Equal(t, data.width, widthDefault)
	require.Equal(t, data.height, heightDefault)
	require.Equal(t, data.backgroundColor, color.Transparent)
	require.Equal(t, data.fontSize, uint(fontSizeDefault)*uint(ratioFontSize))
	require.False(t, data.isColor)
	require.False(t, data.isInverse)
	require.Equal(t, data.curve, curveDefault)
	require.Empty(t, data.ignoreCharacters)
	require.Equal(t, data.charactersPreset, characters)
}

func TestGetOptionByMath(t *testing.T) {
	option := OptionMath{}
	data := getOptionByMath(option)

	require.NotNil(t, data)
	require.Equal(t, data.width, widthDefault)
	require.Equal(t, data.height, heightDefault)
	require.Equal(t, data.backgroundColor, color.Transparent)
	require.Equal(t, data.fontSize, uint(fontSizeDefault)*uint(ratioFontSize))
	require.False(t, data.isColor)
	require.False(t, data.isInverse)
	require.Equal(t, data.curve, curveDefault)
	require.Equal(t, data.mathMin, mathMinDefault)
	require.Equal(t, data.mathMax, mathMaxDefault)
	require.Contains(t, []mathOperator{MathOperatorMinus, MathOperatorPlus}, data.mathOperator)
}

func TestGetOptionByTextWithOption(t *testing.T) {
	option := OptionText{
		Size:             6,
		Width:            200,
		Height:           100,
		FontSize:         18,
		Curve:            3,
		IgnoreCharacters: "90",
		CharactersPreset: "1234567890",
		BackgroundColor:  color.Black,
	}
	opt := getOptionByText(option)

	require.Equal(t, opt.size, uint8(6))
	require.Equal(t, opt.width, uint16(200))
	require.Equal(t, opt.height, uint16(100))
	require.Equal(t, opt.fontSize, uint(18)*uint(ratioFontSize))
	require.Equal(t, opt.curve, uint8(3))
	require.Equal(t, opt.ignoreCharacters, "90")
	require.Equal(t, opt.charactersPreset, "1234567890")
	require.True(t, opt.isColor)
	require.Equal(t, opt.backgroundColor, color.Black)
}

func TestGetOptionByMathWithOption(t *testing.T) {
	option := OptionMath{
		Width:           200,
		Height:          100,
		FontSize:        18,
		Curve:           3,
		BackgroundColor: color.Black,
		MathOperator:    MathOperatorPlus,
		MathMin:         10,
		MathMax:         99,
	}
	opt := getOptionByMath(option)

	require.Equal(t, opt.width, uint16(200))
	require.Equal(t, opt.height, uint16(100))
	require.Equal(t, opt.fontSize, uint(18)*uint(ratioFontSize))
	require.True(t, opt.isColor)
	require.Equal(t, opt.backgroundColor, color.Black)
	require.Equal(t, opt.mathMin, uint8(10))
	require.Equal(t, opt.mathMax, uint16(99))
	require.Equal(t, opt.mathOperator, MathOperatorPlus)
}
