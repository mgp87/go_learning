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

const num = 5

func RestOfVariables() {
	fmt.Println("This is a constant: ", num)

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
