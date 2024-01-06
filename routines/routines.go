package routines

import (
	"fmt"
	"strings"
	"time"
)

func SlowName(name string, c chan bool) {
	letters := strings.Split(name, "")
	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond) // Sleep for 1 second
		fmt.Println(letter)
	}
	c <- true
}
