package exercises

import (
	"fmt"
	"strconv"
)

func GetConditionalResponse(value string) (int, string) {
	num, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("There was an error converting the value to an integer")
		panic(err)
		// return 0, ""
	}
	var result string
	if num > 100 {
		result = "Greater than 100"
	} else {
		result = "Less than 100"
	}
	return num, result
}
