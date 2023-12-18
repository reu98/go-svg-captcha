package captcha

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetOptionByText(t *testing.T) {
	option := OptionText{}
	data := getOptionByText(option)

	require.NotNil(t, data)
	require.Nil(t, data.Size)
	require.Nil(t, data.Width)
	require.Nil(t, data.Height)
	require.Nil(t, data.FontSize)
	require.Nil(t, data.IsColor)
	require.Nil(t, data.IsInverse)
	require.Nil(t, data.Noise)
	require.Nil(t, data.BackgroundColor)
	require.Nil(t, data.Text)
	require.Nil(t, data.IgnoreCharacters)
	require.Nil(t, data.CharactersPreset)
}

func TestGetOptionByMath(t *testing.T) {
	option := OptionMath{}
	data := getOptionByMath(option)

	require.NotNil(t, data)
	require.Nil(t, data.Size)
	require.Nil(t, data.Width)
	require.Nil(t, data.Height)
	require.Nil(t, data.FontSize)
	require.Nil(t, data.IsColor)
	require.Nil(t, data.IsInverse)
	require.Nil(t, data.Noise)
	require.Nil(t, data.BackgroundColor)
	require.Nil(t, data.MathMax)
	require.Nil(t, data.MathMin)
	require.Nil(t, data.MathOperator)
}
