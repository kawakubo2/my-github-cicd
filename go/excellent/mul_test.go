package main

import (
	"testing"
)

func TestMul(t *testing.T) {
	num1 := 4
	num2 := 3
	result := Mul(num1, num2)
	if result != num1 * num2 {
		t.Errorf("expected: %v: actual: %v", num1 * num2, result)
	}
}