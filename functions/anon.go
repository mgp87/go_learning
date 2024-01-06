package functions

import "fmt"

func Calculus(num1 int, num2 int) {
	var num3, num4 int = 2, 5
	addition := func(num1 int, num2 int) int {
		return num1 + num2 + num3 + num4
	}
	fmt.Println(addition(num1, num2))
}
