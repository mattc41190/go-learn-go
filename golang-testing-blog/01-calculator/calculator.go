package main

type calculator struct{}

func (c calculator) Sum(a int, b int) int {
	return a + b
}

func (c calculator) Subtract(a int, b int) int {
	return a - b
}

func (c calculator) Multiply(a int, b int) int {
	return a * b
}

func (c calculator) Divide(a int, b int) int {
	return a / b
}
