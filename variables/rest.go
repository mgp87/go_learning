package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var State bool
var Salary float32
var Date time.Time

const num float32 = 5.3

func RestOfVariables() {
	fmt.Println("This is a constant: ", num)

	const num2 = "5.3"
	fmt.Println("This is another constant: ", num2)

	Name = "Miguel"
	State = true
	Salary = 1500.69
	Date = time.Now()

	fmt.Println("Name: ", Name)
	fmt.Println("Salary: ", Salary)
	fmt.Println("State: ", State)
	fmt.Println("Date: ", Date)
}

func ConvertToText(num int) (bool, string) {
	var text string
	text = strconv.Itoa(num)
	return true, text
}
