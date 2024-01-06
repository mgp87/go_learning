package functions

import "fmt"

func Exponent(num int) { // Recursive function until return condition is met
	if num > 10000000 {
		return
	}
	fmt.Println(num)
	Exponent(num * 2) // Calls itself recursively
}
