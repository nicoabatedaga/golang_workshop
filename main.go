package main

import "fmt"

func main() {

	result, err := division(10, 2)
	fmt.Println("Result:", result)

	result, err = division(10, 0)
	fmt.Println("Error:", err)
}

func division(a, b float64) (float64, error) {
	fmt.Println(fmt.Sprintf("Division of %f by %f", a, b))
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
