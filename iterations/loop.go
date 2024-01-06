package iterations

import "fmt"

func Iter() {
	i := 0
	for i > 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 10; j++ {
		if j == 5 {
			break
		}
		if j == 3 {
			continue
		}
		fmt.Println(j)
	}

	for k := 0; k < 100; k += 5 {
		fmt.Println(k)
	}
}
