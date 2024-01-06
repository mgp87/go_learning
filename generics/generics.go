package generics

import "fmt"

func MyGeneric[T int | float32 | float64 | string](a, b T) T {
	fmt.Println("Init MyGeneric ")
	return a + b
}
