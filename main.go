package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// Use _ to ignore a return value
	// To use _,_ it is used = instead of :=
	answer, err := divide(6, 3)
	if err != nil {
		fmt.Printf("An error occurred %s\n", err.Error())
	} else {
		fmt.Printf("%f\n", answer)
	}
}

func divide(p1, p2 float64) (float64, error) {
	if p2 == 0 {
		return math.NaN(), errors.New("Cannot divide by zero")
	}
	return p1 / p2, nil
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
