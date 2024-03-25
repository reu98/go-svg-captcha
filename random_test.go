package captcha

import (
	"image/color"
	"math/rand"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

const regexColor = "^#(?:[0-9a-fA-F]{3}){1,2}$"

func TestRandomInt(t *testing.T) {
	min := uint8(rand.Uint32())
	max := uint16(rand.Uint32())

	data := randomInt(min, max)

	require.GreaterOrEqual(t, data, uint16(min))
	require.LessOrEqual(t, data, max)
}

func TestGetLightnessWithValidColor(t *testing.T) {
	red := randomInt(1, 255)
	green := randomInt(1, 255)
	blue := randomInt(1, 255)
	backgroundColor := color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: 255}
	opt := &option{
		backgroundColor: backgroundColor,
	}

	data := opt.getLightness()

	require.LessOrEqual(t, data, lightnessDefault)
}

func TestGetMaxColor(t *testing.T) {
	testCase := []struct {
		input    []uint32
		expected uint32
	}{
		{
			input:    []uint32{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			input:    []uint32{99, 256, 0, 10, 100},
			expected: 256,
		},
		{
			input:    []uint32{9, 4, 0, 5, 10},
			expected: 10,
		},
		{
			input:    []uint32{100, 200, 150, 300},
			expected: 300,
		},
	}

	for _, tc := range testCase {
		data := getMaxColor(tc.input...)
		require.Equal(t, data, tc.expected)
	}
}

func TestGetMinColor(t *testing.T) {
	testCase := []struct {
		input    []uint32
		expected uint32
	}{
		{
			input:    []uint32{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			input:    []uint32{99, 256, 1, 10, 100},
			expected: 1,
		},
		{
			input:    []uint32{99, 2, 0, 5, 10},
			expected: 0,
		},
		{
			input:    []uint32{100, 200, 150, 256},
			expected: 100,
		},
	}

	for _, tc := range testCase {
		data := getMinColor(tc.input...)
		require.Equal(t, data, tc.expected)
	}
}

func TestRandomColorWithoutBg(t *testing.T) {
	maxTestCase := 50
	regex := regexp.MustCompile(regexColor)
	opt := &option{
		backgroundColor: color.Transparent,
	}

	for i := 0; i < maxTestCase; i++ {
		data := opt.randomColor()
		require.Regexp(t, regex, data)
	}
}

func TestRandomColorWithBg(t *testing.T) {
	maxTestCase := 50
	regex := regexp.MustCompile(regexColor)
	opt := &option{
		backgroundColor: color.Black,
	}

	for i := 0; i < maxTestCase; i++ {
		data := opt.randomColor()
		require.Regexp(t, regex, data)
	}
}

func TestRemoveCharacters(t *testing.T) {
	testCase := []struct {
		characters  string
		ignoreChars string
		expected    string
	}{
		{
			characters:  "ABC",
			ignoreChars: "A",
			expected:    "BC",
		},
		{
			characters:  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
			ignoreChars: "0123456789",
			expected:    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		},
		{
			characters:  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
			ignoreChars: "AaBbOd76",
			expected:    "CDEFGHIJKLMNPQRSTUVWXYZcefghijklmnopqrstuvwxyz01234589",
		},
	}

	for _, tc := range testCase {
		data := removeCharacters(tc.characters, tc.ignoreChars)

		require.Equal(t, data, tc.expected)
	}
}

func TestRandomGreyColor(t *testing.T) {
	min := uint8(rand.Uint32())
	max := uint8(rand.Uint32())
	if min > max {
		swap(&min, &max)
	}

	data := randomGreyColor(&min, &max)

	require.True(t, isGreyColor(data))
}

func TestRandomText(t *testing.T) {
	opt := getOptionByText(OptionText{})

	data := opt.randomText()

	require.Len(t, data, int(sizeDefault))
}

func TestRandomTextWithSize(t *testing.T) {
	size := uint8(rand.Intn(len(characters)))
	opt := getOptionByText(OptionText{
		Size: size,
	})

	data := opt.randomText()

	require.Len(t, data, int(size))
}

func TestRandomTextWithPreset(t *testing.T) {
	charactersPreset := "0123456789"
	option := &option{
		charactersPreset: charactersPreset,
	}

	data := option.randomText()
	for _, char := range data {
		require.Contains(t, charactersPreset, string(char))
	}
}

func TestRandomTextWithIgnoreCharacters(t *testing.T) {
	charactersPreset := "0123456789"
	ignoreCharacters := "095"
	option := &option{
		charactersPreset: charactersPreset,
		ignoreCharacters: ignoreCharacters,
	}

	data := option.randomText()
	for _, char := range data {
		require.Contains(t, charactersPreset, string(char))
		require.NotContains(t, ignoreCharacters, string(char))
	}
}

func swap(a, b *uint8) {
	temp := *a
	*a = *b
	*b = temp
}

func isGreyColor(color string) bool {
	if len(color) != 7 || color[0] != '#' {
		return false
	}

	red, _ := strconv.ParseUint(color[1:3], 16, 8)
	green, _ := strconv.ParseUint(color[3:5], 16, 8)
	blue, _ := strconv.ParseUint(color[5:7], 16, 8)

	var threshold uint64 = 16
	return absDiff(red, green) <= threshold &&
		absDiff(green, blue) <= threshold &&
		absDiff(blue, red) <= threshold
}

func absDiff(a, b uint64) uint64 {
	if a > b {
		return a - b
	}
	return b - a
}
