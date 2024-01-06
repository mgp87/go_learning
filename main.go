package main

import (
	"fmt"
	"runtime"

	"github.com/mgp87/go_learning/exercises"
	"github.com/mgp87/go_learning/files"
	"github.com/mgp87/go_learning/functions"
	"github.com/mgp87/go_learning/iterations"
	"github.com/mgp87/go_learning/keyboard"
	"github.com/mgp87/go_learning/variables"
)

func main() {
	variables.ShowIntegers()
	variables.RestOfVariables()
	result, value := variables.ConvertToText(1776)
	fmt.Println("Result: ", result)
	fmt.Println("Value: ", value)

	fmt.Println(runtime.GOOS)
	if os := runtime.GOOS; os == "linux" || os == "OS X" || os == "darwin" {
		fmt.Println("Not Windows")
	} else {
		fmt.Println("Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("Darwin (OS X)")
	default:
		fmt.Printf("%s \n", os)
	}

	num, conclusion := exercises.GetConditionalResponse("150")
	fmt.Println("Number: ", num)
	fmt.Println("Conclusion: ", conclusion)

	keyboard.GetNumbers()
	iterations.Iter()

	// fmt.Println(exercises.MulTable())

	files.SaveTable()
	files.AddTable()
	files.ReadFile()

	functions.Calculus(3, 8)
	functions.CallClosure()
	functions.Exponent(3)
}
