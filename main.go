package main

import "fmt"

func main() {
	fmt.Printf("%f", divide(6, 3))
}

func divide(p1, p2 float64) float64 {
	return p1 / p2
}

func add(p1, p2 float64) float64 {
	return p1 + p2
}

func subtract(p1, p2 float64) float64 {
	return p1 - p2
}

func multiply(p1, p2 float64) float64 {
	return p1 * p2
}
