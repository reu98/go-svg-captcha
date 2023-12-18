package captcha

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomOffset(t *testing.T) {
	data := randomOffset()

	require.LessOrEqual(t, data, lightnessDefault)
}

func TestRandomOperation(t *testing.T) {
	var dataType matchOperator
	data := randomOperation()

	require.IsType(t, dataType, data)
	require.Contains(t, []matchOperator{MathOperatorMinus, MathOperatorPlus}, data)
}

func TestDrawText(t *testing.T) {
	option := option{}
	text := option.randomText()
	data, err := option.drawText(text)

	require.NoError(t, err)
	require.NotEmpty(t, data)
}

func TestDrawTextWithoutDefault(t *testing.T) {
	width := uint16(rand.Uint32())
	height := uint16(rand.Uint32())
	isColor := true
	isInverse := true
	option := option{
		Width:     &width,
		Height:    &height,
		IsColor:   &isColor,
		IsInverse: &isInverse,
	}
	text := option.randomText()
	data, err := option.drawText(text)

	require.NoError(t, err)
	require.NotEmpty(t, data)
}
