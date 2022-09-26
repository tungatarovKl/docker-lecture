package main

import (
	"errors"
)

func calculate(a, b int, operation string) (int, error) {
	switch operation {
	case "+":
		return sum(a, b), nil
	case "-":
		return subtract(a, b), nil
	case "*":
		return multiply(a, b), nil
	case "/":
		return divide(a, b)
	default:
		return 0, errors.New("не разрешенная операция")
	}
}

func sum(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль невозможна")
	}
	
	return a / b, nil
}

func multiply(a, b int) int {
	return a * b
}
