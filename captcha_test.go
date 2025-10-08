package captcha

import (
	"image/color"
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
	minMath := uint8(10)
	maxMath := uint16(100)
	option := OptionMath{
		MathMin: minMath,
		MathMax: maxMath,
	}
	data, err := CreateByMath(option)

	require.NoError(t, err)
	require.NotNil(t, data)
	require.IsType(t, dataType, *data)
	require.NotEmpty(t, (*data).Data)

	result, err := strconv.Atoi(data.Text)
	require.NoError(t, err)

	// For subtraction: minMath - maxMath (could be negative)
	// For addition: maxMath + maxMath (maximum possible)
	require.GreaterOrEqual(t, result, int(minMath)-int(maxMath))
	require.LessOrEqual(t, result, int(maxMath)+int(maxMath))
}

func TestCreateCaptcha(t *testing.T) {
	opt := getOptionByText(OptionText{
		BackgroundColor: color.Black,
	})
	text := opt.randomText()
	capt, err := opt.createCaptcha(text)

	require.NotEmpty(t, capt)
	require.NoError(t, err)
}

func TestCreateCaptchaWithText(t *testing.T) {
	opt := getOptionByText(OptionText{
		Text: "abcd",
	})
	capt, err := opt.createCaptcha(opt.text)

	require.NotEmpty(t, capt)
	require.NoError(t, err)
}
