package main

import "errors"

// ServiceMetricMath Define a service interface
type ServiceMetricMath interface {

	// Add calculate a+b
	Add(a, b int) int

	// Subtract calculate a-b
	Subtract(a, b int) int

	// Multiply calculate a*b
	Multiply(a, b int) int

	// Divide calculate a/b
	Divide(a, b int) (int, error)
}

//ArithmeticServiceMath implement ServiceMetricMath interface
type ArithmeticServiceMath struct {
}

// Add implement Add method
func (s ArithmeticServiceMath) Add(a, b int) int {
	return a + b
}

// Subtract implement Subtract method
func (s ArithmeticServiceMath) Subtract(a, b int) int {
	return a - b
}

// Multiply implement Multiply method
func (s ArithmeticServiceMath) Multiply(a, b int) int {
	return a * b
}

// Divide implement Divide method
func (s ArithmeticServiceMath) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("the dividend can not be zero!")
	}

	return a / b, nil
}
