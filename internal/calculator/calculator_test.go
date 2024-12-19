package calculator

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		shouldErr  bool
	}{
		{"2+2*2", 6, false},
		{"(1 + 2) * 3 - 4 / 2", 7, false},
		{"10 / (5 - 5)", 0, true},
		{"2 + a", 0, true},
		{"(2+3", 5, false},
		{"", 0, true},
		{"3 + 4 * 2 / (1 - 5)^2", 0, true},
	}

	for _, test := range tests {
		result, err := Calc(test.expression)
		if test.shouldErr {
			if err == nil {
				t.Errorf("Ожидалась ошибка для выражения %q, но ошибки не было", test.expression)
			}
		} else {
			if err != nil {
				t.Errorf("Не ожидалась ошибка для выражения %q, но получена: %v", test.expression, err)
			}
			if result != test.expected {
				t.Errorf("Для выражения %q ожидается %v, получено %v", test.expression, test.expected, result)
			}
		}
	}
}
