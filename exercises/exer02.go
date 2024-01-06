package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MulTable() string {
	var result string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a number: ")
	if scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Incorrect data: " + err.Error())
		}
		for i := range [11]int{} {
			result += fmt.Sprintf("%d x %d = %d\n", num, i, num*i)
		}
		return result
	}
	return ""
}
