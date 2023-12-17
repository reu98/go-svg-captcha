package captcha

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateMathOperationPlus(t *testing.T) {
	operandLeft := uint16(rand.Uint32())
	operandRight := uint16(rand.Uint32())
	equationResult := fmt.Sprintf("%v + %v", operandLeft, operandRight)

	data := generateMathOperationPlus(operandLeft, operandRight)

	result, err := strconv.Atoi(data.Text)
	require.NoError(t, err)

	require.NotNil(t, data)
	require.Equal(t, result, int(operandLeft)+int(operandRight))
	require.Equal(t, data.Equation, equationResult)
}

func TestGenerateMathOperationMinus(t *testing.T) {
	operandLeft := uint16(rand.Uint32())
	operandRight := uint16(rand.Uint32())
	equationResult := fmt.Sprintf("%v - %v", operandLeft, operandRight)

	data := generateMathOperationMinus(operandLeft, operandRight)

	result, err := strconv.Atoi(data.Text)
	require.NoError(t, err)

	require.NotNil(t, data)
	require.Equal(t, result, int(operandLeft)-int(operandRight))
	require.Equal(t, data.Equation, equationResult)
}

func TestGenerateMathOperationByPlus(t *testing.T) {
	operator := MathOperatorPlus
	min := uint8(rand.Uint32())
	max := uint16(rand.Uint32())

	data := generateMathOperation(&min, &max, &operator)

	require.NotNil(t, data)

	result, err := strconv.Atoi(data.Text)
	require.NoError(t, err)

	require.Greater(t, result, 0)
	require.LessOrEqual(t, result, int(max*2))
}

func TestGenerateMathOperationByMinus(t *testing.T) {
	operator := MathOperatorMinus
	min := uint8(rand.Uint32())
	max := uint16(rand.Uint32())

	data := generateMathOperation(&min, &max, &operator)
	require.NotNil(t, data)
	fmt.Printf("%v - %v - %v\n", min, max, data.Text)
	result, err := strconv.Atoi(data.Text)
	require.NoError(t, err)

	require.GreaterOrEqual(t, result, int(max)*-2)
	require.LessOrEqual(t, result, int(max))
}

func TestGenerateMathOperation(t *testing.T) {
	data := generateMathOperation(nil, nil, nil)

	require.NotNil(t, data)

	result, err := strconv.Atoi(data.Text)
	require.NoError(t, err)

	require.GreaterOrEqual(t, result, -2*int(mathMaxDefault))
	require.LessOrEqual(t, result, int(mathMaxDefault*2))
}
