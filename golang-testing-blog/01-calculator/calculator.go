package main

type calculator struct{}

func (c calculator) Sum(a int, b int) int {
	return a + b
}
