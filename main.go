package main

import (
	"fmt"
	"functions/simplemath"
	"net/http"
	"strings"
)

type MathExpr = string

const (
	AddExpr       = MathExpr("add")
	SubtractdExpr = MathExpr("subtact")
	MultiplyExpr  = MathExpr("multiply")
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

	sv := simplemath.NewSemanticVersion(1, 2, 3)
	sv.IncrementMajor()
	println(sv.String())

	var tripper http.RoundTripper = RoundTripCounter{}
	r, _ := http.NewRequest(http.MethodGet, "http://pluralsight.com", strings.NewReader("test call"))
	_, _ = tripper.RoundTrip(r)

	a := func(name string) string {
		fmt.Printf("My first %s function\n", name)
		return name
	}

	value := a("miguel")
	println(value)

	addExpr := mathExpression(MultiplyExpr)
	println(addExpr(2, 3))
}

type RoundTripCounter struct {
	count int
}

func (rt RoundTripCounter) RoundTrip(*http.Request) (*http.Response, error) {
	rt.count += 1
	return nil, nil
}

func mathExpression(expr MathExpr) func(float64, float64) float64 {
	// return func(f float64, f2 float64) float64 {
	// 	return f + f2
	// }
	switch expr {
	case AddExpr:
		return simplemath.Add
	case SubtractdExpr:
		return simplemath.Subtract
	case MultiplyExpr:
		return simplemath.Multiply
	default:
		return func(f float64, f2 float64) float64 {
			return 0
		}
	}
}
