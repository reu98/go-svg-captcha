package captcha

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateByText(t *testing.T) {
	var dataType Result
	option := OptionText{}
	data, err := CreateByText(option)

	require.NoError(t, err)
	require.NotNil(t, data)
	require.IsType(t, dataType, *data)
	require.Len(t, (*data).Text, int(sizeDefault))
	require.NotEmpty(t, (*data).Data)
}

func TestCreateByTextByOption(t *testing.T) {
	var dataType Result
	text := "ABCD"
	option := OptionText{
		Text: &text,
	}
	data, err := CreateByText(option)

	require.NoError(t, err)
	require.NotNil(t, data)
	require.IsType(t, dataType, *data)
	require.Len(t, (*data).Text, int(sizeDefault))
	require.NotEmpty(t, (*data).Data)
}

func TestCreateByMath(t *testing.T) {
	var dataType Result
	maxMath := 18
	minMath := -8
	option := OptionMath{}
	data, err := CreateByMath(option)

	require.NoError(t, err)
	require.NotNil(t, data)
	require.IsType(t, dataType, *data)
	require.NotEmpty(t, (*data).Data)

	result, err := strconv.Atoi((*data).Text)
	require.NoError(t, err)

	require.LessOrEqual(t, result, maxMath)
	require.GreaterOrEqual(t, result, minMath)
}

func TestCreateByMathByOption(t *testing.T) {
	var dataType Result
	minMath := uint8(rand.Uint32())
	maxMath := uint16(minMath) + uint16(rand.Uint32())
	option := OptionMath{
		MathMin: &minMath,
		MathMax: &maxMath,
	}
	data, err := CreateByMath(option)

	require.NoError(t, err)
	require.NotNil(t, data)
	require.IsType(t, dataType, *data)
	require.NotEmpty(t, (*data).Data)

	result, err := strconv.Atoi((*data).Text)
	require.NoError(t, err)

	require.GreaterOrEqual(t, result, int(minMath)-int(maxMath))
	require.LessOrEqual(t, result, int(maxMath)*2)
}
