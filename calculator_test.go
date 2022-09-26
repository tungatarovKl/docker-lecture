package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result := sum(1, 2)
	
	if !assert.Equal(t, 3, result) {
		t.Errorf("sum: Ожидалось %d, получено: %d", 3, result)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(4, 2)
	
	if !assert.Equal(t, 2, result) {
		t.Errorf("sum: Ожидалось %d, получено: %d", 2, result)
	}
}

func TestMultiply(t *testing.T) {
	result := multiply(5, 2)
	
	if !assert.Equal(t, 10, result) {
		t.Errorf("sum: Ожидалось %d, получено: %d", 10, result)
	}
}

func TestDivide(t *testing.T) {
	result, err := divide(6, 2)
	if err != nil {
		t.Errorf("Неожиданная ошибка при тестирование. Ожидалось %d, получена ошибка", 3)
	}
	
	if !assert.Equal(t, 3, result) {
		t.Errorf("sum: Ожидалось %d, получено: %d", 3, result)
	}
}
