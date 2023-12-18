package captcha

import (
	"fmt"
	"strconv"
)

type resultMath struct {
	Text     string
	Equation string
}

func generateMathOperation(min *uint8, max *uint16, operator *matchOperator) *resultMath {
	mathMin := mathMinDefault
	if min != nil {
		mathMin = *min
	}

	mathMax := mathMaxDefault
	if max != nil {
		mathMax = *max
	}

	mathOperator := MathOperatorPlus
	if operator != nil {
		mathOperator = *operator
	}

	operandLeft := randomInt(mathMin, mathMax)
	operandRight := randomInt(mathMin, mathMax)

	if mathOperator == MathOperatorPlus {
		return generateMathOperationPlus(operandLeft, operandRight)
	}

	return generateMathOperationMinus(operandLeft, operandRight)
}

func generateMathOperationPlus(operandLeft, operandRight uint16) *resultMath {
	result := strconv.Itoa(int(operandLeft) + int(operandRight))
	equation := fmt.Sprintf("%v + %v", operandLeft, operandRight)

	return &resultMath{
		Text:     result,
		Equation: equation,
	}
}

func generateMathOperationMinus(operandLeft, operandRight uint16) *resultMath {
	result := strconv.Itoa(int(operandLeft) - int(operandRight))
	equation := fmt.Sprintf("%v - %v", operandLeft, operandRight)

	return &resultMath{
		Text:     result,
		Equation: equation,
	}
}
