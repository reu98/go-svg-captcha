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
