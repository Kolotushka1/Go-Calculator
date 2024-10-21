package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	tokens := tokenize(expression)
	if tokens == nil {
		return 0, fmt.Errorf("некорректное выражение")
	}
	result, err := parse(tokens)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func tokenize(expr string) []string {
	var tokens []string
	var number strings.Builder

	for _, ch := range expr {
		if ch == ' ' {
			continue
		}
		if (ch >= '0' && ch <= '9') || ch == '.' {
			number.WriteRune(ch)
		} else {
			if number.Len() > 0 {
				tokens = append(tokens, number.String())
				number.Reset()
			}
			if strings.ContainsRune("+-*/()", ch) {
				tokens = append(tokens, string(ch))
			} else {
				return nil
			}
		}
	}
	if number.Len() > 0 {
		tokens = append(tokens, number.String())
	}
	return tokens
}

func parse(tokens []string) (float64, error) {
	result, _, err := parseSub(tokens, 0)
	return result, err
}

func parseSub(tokens []string, index int) (float64, int, error) {
	var values []float64
	var operators []string

	for index < len(tokens) {
		token := tokens[index]

		if token == "(" {
			val, newIndex, err := parseSub(tokens, index+1)
			if err != nil {
				return 0, index, err
			}
			values = append(values, val)
			index = newIndex
		} else if token == ")" {
			break
		} else if isOperator(token) {
			for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(token) {
				var err error
				values, operators, err = applyOperator(values, operators)
				if err != nil {
					return 0, index, err
				}
			}
			operators = append(operators, token)
		} else {
			val, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, index, fmt.Errorf("некорректное число: %s", token)
			}
			values = append(values, val)
		}
		index++
	}

	for len(operators) > 0 {
		var err error
		values, operators, err = applyOperator(values, operators)
		if err != nil {
			return 0, index, err
		}
	}

	if len(values) != 1 {
		return 0, index, fmt.Errorf("некорректное выражение")
	}

	return values[0], index, nil
}

func isOperator(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/"
}

func precedence(op string) int {
	if op == "+" || op == "-" {
		return 1
	}
	if op == "*" || op == "/" {
		return 2
	}
	return 0
}

func applyOperator(values []float64, operators []string) ([]float64, []string, error) {
	if len(values) < 2 || len(operators) == 0 {
		return values, operators, fmt.Errorf("некорректное выражение")
	}

	right := values[len(values)-1]
	left := values[len(values)-2]
	values = values[:len(values)-2]

	op := operators[len(operators)-1]
	operators = operators[:len(operators)-1]

	var result float64
	switch op {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		if right == 0 {
			return values, operators, fmt.Errorf("деление на ноль")
		}
		result = left / right
	default:
		return values, operators, fmt.Errorf("неизвестный оператор: %s", op)
	}

	values = append(values, result)
	return values, operators, nil
}

func main() {
	example := "(1 + 2) * 3 - 4 / 2"
	result, err := Calc(example)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}
