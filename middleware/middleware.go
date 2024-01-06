package middleware

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func MyMiddleware() {
	fmt.Println("Init Middleware")
	result := middlewareOps(add)(3, 5)
	fmt.Println("Result: ", result)
	result = middlewareOps(subtract)(3, 5)
	fmt.Println("Result: ", result)
	result = middlewareOps(multiply)(3, 5)
	fmt.Println("Result: ", result)
}

func middlewareOps(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Init Middleware Ops")
		return f(a, b)
	}
}
