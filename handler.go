package lab2

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
)

type ComputeHandler struct {
	input  string
	output io.Writer
}

func NewComputeHandler(input string, output io.Writer) *ComputeHandler {
	return &ComputeHandler{input: input, output: output}
}

func (ch *ComputeHandler) Compute() error {
	expression := strings.Fields(ch.input)
	if len(expression) == 0 {
		return errors.New("відсутні дані для обробки")
	}

	result, err := PostfixToInfix(strings.Join(expression, " "))
	if err != nil {
		return err
	}

	_, err = ch.output.Write([]byte(result))
	if err != nil {
		return err
	}

	return nil
}

func ParseInput(expressionFlag, fileFlag string) (string, error) {
	if expressionFlag != "" {
		return expressionFlag, nil
	} else if fileFlag != "" {
		data, err := ioutil.ReadFile(fileFlag)
		if err != nil {
			return "", err
		}
		return string(data), nil
	} else {
		return "", errors.New("не вказаний вираз або файл з виразом")
	}
}
