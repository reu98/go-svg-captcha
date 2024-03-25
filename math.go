package captcha

import (
	"fmt"
	"strconv"
)

type resultMath struct {
	Text     string
	Equation string
}

func (opt *option) generateMathOperation() *resultMath {
	operandLeft := randomInt(opt.mathMin, opt.mathMax)
	operandRight := randomInt(opt.mathMin, opt.mathMax)

	if opt.mathOperator == MathOperatorPlus {
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
