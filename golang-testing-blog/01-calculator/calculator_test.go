package main

import "testing"

func TestCalculatorSum(t *testing.T) {
	c := calculator{}
	sum := c.Sum(2, 2)
	if sum != 4 {
		t.Errorf("Expected 4, but got %d", sum)
	}
}
