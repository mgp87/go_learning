package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var legend string
var err error

func GetNumbers() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the first number: ")
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Incorrect data: " + err.Error())
		}
	}
	fmt.Println("Enter the second number: ")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Incorrect data: " + err.Error())
		}
	}

	fmt.Println("Enter the legend: ")
	if scanner.Scan() {
		legend = scanner.Text()
	}

	fmt.Println("The result is: ", num1*num2, legend)
}
