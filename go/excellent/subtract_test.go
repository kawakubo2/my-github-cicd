package main

import (
	"testing"
)

func TestSubtract(t *testing.T) {
	num1 := 12
	num2 := 7
	result := Sub(num1, num2)
	if result != num1 - num2 {
		t.Errorf("%v - %v = %v", num1, num2, result)
	}
}