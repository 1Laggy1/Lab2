// handler_test.go

package lab2

import (
	"bytes"
	"errors"
	"testing"
)

func TestComputeHandler_Compute(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedOut string
		expectedErr error
	}{
		{
			name:        "Valid expression",
			input:       "2 3 +",
			expectedOut: "(2+3)",
			expectedErr: nil,
		},
		{
			name:        "Invalid expression",
			input:       "2 3 + +",
			expectedOut: "",
			expectedErr: errors.New("invalid expression: insufficient operands"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var output bytes.Buffer

			handler := NewComputeHandler(test.input, &output)
			err := handler.Compute()

			if err != nil && test.expectedErr == nil {
				t.Errorf("unexpected error: %v", err)
			} else if err == nil && test.expectedErr != nil {
				t.Errorf("expected error: %v, got nil", test.expectedErr)
			} else if err != nil && test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
				t.Errorf("expected error: %v, got: %v", test.expectedErr, err)
			}

			if output.String() != test.expectedOut {
				t.Errorf("expected output: %s, got: %s", test.expectedOut, output.String())
			}
		})
	}
}
