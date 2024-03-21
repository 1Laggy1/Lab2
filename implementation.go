//implementation.go

package lab2

import (
	"fmt"
	"strings"
)

func PostfixToInfix(expression string) (string, error) {
	stack := []string{}

	tokens := strings.Fields(expression)

	for _, token := range tokens {
		if isOperator(token) {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid expression: insufficient operands")
			}

			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			stack = append(stack, "("+operand1+token+operand2+")")
		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid expression: invalid number of operands")
	}

	return stack[0], nil
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
}

func main() {
	postfixExpression := "10 2 8 * 4 / + 3 -"
	infixExpression, err := PostfixToInfix(postfixExpression)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Infix expression:", infixExpression)
}
