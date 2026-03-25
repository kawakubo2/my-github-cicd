package main

import(
	"testing"
)

func TestAdd(t *testing.T) {
	num1 := 8
	num2 := 5
	result = Add(num1, num2)
	if result != num1 + num2 {
		t.Errorf("expected: %v. actual: %v", num1 + num2, result)
	}
}