package main

import (
	"fmt"

	"github.com/mgp87/go_learning/variables"
)

func main() {
	variables.ShowIntegers()
	variables.RestOfVariables()
	result, value := variables.ConvertToText(1776)
	fmt.Println("Result: ", result)
	fmt.Println("Value: ", value)
}
