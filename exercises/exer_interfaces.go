package exercises

import (
	"fmt"

	"github.com/mgp87/go_learning/interfaces"
)

func BreathingHumans(human interfaces.Human) {
	human.Breathe()
	fmt.Printf("I am a %s and I am breathing\n", human.Sex())
}
