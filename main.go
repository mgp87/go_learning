package main

import (
	"fmt"
	"runtime"

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
}
