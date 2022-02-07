package main

// ServiceMetricMath Define a service interface
type ServiceMetricString interface {

	// Add calculate a+b
	Add(a, b string) string
}

//ArithmeticServiceMath implement ServiceMetricMath interface
type ArithmeticServiceString struct {
}

// Add implement Add method
func (s ArithmeticServiceString) Add(a, b string) string {
	return a + b
}
