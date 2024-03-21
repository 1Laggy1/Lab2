package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix(t *testing.T) {
	cases := []struct {
		expression string
		expected   string
		err        bool
	}{
		{"10 2 8 * 4 / + 3 -", "((10+((2*8)/4))-3)", false},
		{"4 2 - 3 * 5 +", "(((4-2)*3)+5)", false},
		{"2 3 + 5 2 - *", "((2+3)*(5-2))", false},
		{"", "", true},
		{"1 2 3 +", "", true},
	}

	for _, tc := range cases {
		result, err := PostfixToInfix(tc.expression)
		if tc.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		}
	}
}
