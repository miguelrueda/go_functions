package main

import (
	"errors"
	"fmt"
	"functions/simplemath"
	"io"
	"math"
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

	fmt.Printf("%f", double(3, 2, mathExpression(AddExpr)))

	// p2 := powerOfTwo()
	// value = p2()

	// println(value)
	// value = p2()
	// println(value)

	var funcs []func() int64
	for i := 0; i < 10; i++ {
		cleanI := i
		funcs = append(funcs, func() int64 {
			return int64(math.Pow(float64(cleanI), 2))
		})
	}

	for _, f := range funcs {
		println(f())
	}

	// ReadSomethingBad()
	if err := ReadFullFile(); err != nil {
		fmt.Printf("something bad occurred %s", err)
	}
}

func ReadFullFile() (err error) {
	var r io.ReadCloser = &SimpleReader{}
	defer func() {
		_ = r.Close()
		if p := recover(); p == errCatastrophicReader {
			println(p)
			err = errors.New("A panic occurred but it is ok")
		} else if p != nil {
			panic("an unexpected error occurred and we do not want to recover")
		}
	}()

	defer func() {
		println("before for-loop")
	}()
	for {
		value, readerErr := r.Read([]byte("text that does nothing"))
		if readerErr == io.EOF {
			println("finished reading file, breaking out of the loop")
			break
		} else if readerErr != nil {
			return readerErr
		}
		println(value)
	}
	defer func() {
		println("after for-loop")
	}()
	return nil
}

func ReadSomethingBad() error {
	var r io.Reader = BadReader{errors.New("My nonsense reader")}
	_, err := r.Read([]byte("test something"))
	if err != nil {
		fmt.Printf("an error occured %s", err)
		return err
	}

	return nil
}

type BadReader struct {
	err error
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}

type SimpleReader struct {
	count int
}

var errCatastrophicReader = errors.New("something catastrphic occurred in the reader")

func (br *SimpleReader) Read(p []byte) (n int, err error) {
	// println(br.count)
	if br.count == 2 {
		panic(errors.New("another error"))
	}
	if br.count > 3 {
		return 0, io.EOF
	}
	br.count += 1
	return br.count, nil
}

func (br *SimpleReader) Close() error {
	println("Closing reader")
	return nil
}

type RoundTripCounter struct {
	count int
}

func (rt RoundTripCounter) RoundTrip(*http.Request) (*http.Response, error) {
	rt.count++
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
		// return func(f float64, f2 float64) float64 {
		// 	return 0
		// }
		panic("an invalid math expression was used")
	}
}

func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return 2 * mathExpr(f1, f2)
}

func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
}
