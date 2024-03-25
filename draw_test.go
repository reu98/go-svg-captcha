package captcha

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomOffset(t *testing.T) {
	data := randomOffset()

	require.LessOrEqual(t, data, float32(lightnessDefault))
}

func TestRandomOperation(t *testing.T) {
	var dataType mathOperator
	data := randomOperation()

	require.IsType(t, dataType, data)
	require.Contains(t, []mathOperator{MathOperatorMinus, MathOperatorPlus}, data)
}

func TestDrawText(t *testing.T) {
	opt := getOptionByText(OptionText{})
	text := opt.randomText()
	data, err := opt.drawText(text)

	require.NoError(t, err)
	require.NotEmpty(t, data)
}

func TestDrawTextWithoutDefault(t *testing.T) {
	width := uint16(rand.Uint32())
	height := uint16(rand.Uint32())
	opt := getOptionByText(OptionText{
		Width:     width,
		Height:    height,
		IsColor:   true,
		IsInverse: true,
	})
	text := opt.randomText()
	data, err := opt.drawText(text)

	require.NoError(t, err)
	require.NotEmpty(t, data)
}
