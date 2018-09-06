package main

import "testing"

func TestCalculatorSum(t *testing.T) {
	c := calculator{}
	sum := c.Sum(2, 2)
	if sum != 4 {
		t.Errorf("Expected 4, but got %d", sum)
	}
}

func TestCalculatorSubtract(t *testing.T) {
	c := calculator{}
	diff := c.Subtract(10, 5)
	if diff != 5 {
		t.Errorf("Expected 5, but got %d", diff)
	}
}

func TestCalculatorMultiply(t *testing.T) {
	c := calculator{}
	type timesTable struct {
		firstNum  int
		secondNum int
		product   int
	}
	tables := []timesTable{
		{2, 3, 6},
		{4, 4, 16},
		{1, 2, 2},
	}

	for _, table := range tables {
		product := c.Multiply(table.firstNum, table.secondNum)
		if product != table.product {
			t.Errorf("FAIL")
		}

	}
}

func TestCalculatorDivide(t *testing.T) {
	c := calculator{}

	type divisionTable struct {
		firstNum  int
		secondNum int
		quotient  int
	}

	tables := []divisionTable{
		{4, 2, 2},
		{6, 3, 2},
		{12, 2, 6},
	}

	for _, table := range tables {
		quotient := c.Divide(table.firstNum, table.secondNum)
		if quotient != table.quotient {
			t.Errorf("FAIL")
		}
	}

}
