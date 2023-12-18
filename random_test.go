package captcha

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

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
	hex := fmt.Sprintf("#%02x%02x%02x", red, green, blue)

	data := getLightness(hex)

	require.LessOrEqual(t, data, lightnessDefault)
}

func TestGetLightnessWithValidShortColor(t *testing.T) {
	hex := "#DDD"
	data := getLightness(hex)

	require.LessOrEqual(t, data, lightnessDefault)
}

func TestGetLightnessWithInvalidColor(t *testing.T) {
	hex := "#GGGG"

	data := getLightness(hex)

	require.Equal(t, data, lightnessDefault)
}

func TestGetMax(t *testing.T) {
	testCase := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			input:    []int{99, 299, 0, -10, -100},
			expected: 299,
		},
		{
			input:    []int{-10, -5, 0, 5, 10},
			expected: 10,
		},
		{
			input:    []int{100, 200, 150, 300},
			expected: 300,
		},
	}

	for _, tc := range testCase {
		data := getMax(tc.input...)
		require.Equal(t, data, tc.expected)
	}
}

func TestGetMin(t *testing.T) {
	testCase := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			input:    []int{99, 299, 0, -10, -100},
			expected: -100,
		},
		{
			input:    []int{-10, -5, 0, 5, 10},
			expected: -10,
		},
		{
			input:    []int{100, 200, 150, 300},
			expected: 100,
		},
	}

	for _, tc := range testCase {
		data := getMin(tc.input...)
		require.Equal(t, data, tc.expected)
	}
}

func TestRandomColorWithoutBg(t *testing.T) {
	index := 1
	maxTestCase := 50
	regex := regexp.MustCompile(regexColor)

	for ; index <= maxTestCase; index++ {
		data := randomColor(nil)
		require.Regexp(t, regex, data)
	}
}

func TestRandomColorWithBg(t *testing.T) {
	bgColor := "#000"
	data := randomColor(&bgColor)

	regex := regexp.MustCompile(regexColor)

	require.Regexp(t, regex, data)
}

func TestRandomColorWithInvalidBg(t *testing.T) {
	bgColor := "#ggg"
	data := randomColor(&bgColor)

	regex := regexp.MustCompile(regexColor)

	require.Regexp(t, regex, data)
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
	option := option{}

	data := option.randomText()

	require.Len(t, data, int(sizeDefault))
}

func TestRandomTextWithSize(t *testing.T) {
	size := uint8(rand.Intn(len(characters)))
	option := option{
		Size: &size,
	}

	data := option.randomText()

	require.Len(t, data, int(size))
}

func TestRandomTextWithPreset(t *testing.T) {
	charactersPreset := "0123456789"
	option := option{
		CharactersPreset: &charactersPreset,
	}

	data := option.randomText()
	for _, char := range data {
		require.Contains(t, charactersPreset, string(char))
	}
}

func TestRandomTextWithIgnoreCharacters(t *testing.T) {
	charactersPreset := "0123456789"
	ignoreCharacters := "095"
	option := option{
		CharactersPreset: &charactersPreset,
		IgnoreCharacters: &ignoreCharacters,
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
