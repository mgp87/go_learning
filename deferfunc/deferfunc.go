package deferfunc

import (
	"fmt"
	"log"
)

func CheckDefer() {
	fmt.Println("First Defer message")
	defer fmt.Println("Final Defer message") // This will be executed at the end of the function regardless of where it is placed
	fmt.Println("Second Defer message")
}

func PanicExample() {
	defer func() {
		reco := recover() // Recover from a panic, always use it with defer; if there is no panic, it will return nil
		if reco != nil {
			log.Fatalf("Recover from Panic - Error: %v", reco) // Fatal will stop the execution of the program exiting with status 1 (os.Exit(1))
		}
	}()
	a := 1
	if a == 1 {
		panic("Something went wrong")
	}
}
