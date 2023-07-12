package handlers

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func evaluateExpression(strExpression string) (string, error) {
	expression, err := govaluate.NewEvaluableExpression(strExpression)
	if err != nil {
		return "", err
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		return "", err
	}

	resultStr := fmt.Sprintf("%f", result)

	length := len(resultStr) - 1
	for i := length; i > length-6; i-- {
		if resultStr[i] != '0' {
			break
		}
		resultStr = resultStr[:i]
	}

	if resultStr[len(resultStr)-1] == '.' {
		resultStr = resultStr[:len(resultStr)-1]
	}

	return resultStr, nil
}
