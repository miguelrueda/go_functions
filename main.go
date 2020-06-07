package main

import (
	"fmt"
	"simplemath"
)

func main() {
	// Use _ to ignore a return value
	// To use _,_ it is used = instead of :=
	answer, err := simplemath.Divide(6, 3)
	if err != nil {
		fmt.Printf("An error occurred %s\n", err.Error())
	} else {
		fmt.Printf("%f\n", answer)
	}

	numbers := []float64{12.2, 14, 16, 22.4}
	total := simplemath.Sum(numbers...)
	fmt.Printf("Total of sum %f\n", total)
}
