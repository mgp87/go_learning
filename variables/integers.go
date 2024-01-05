package variables

import "fmt"

func ShowIntegers() {
	var a int
	b := int32(2)
	c := int64(100)

	fmt.Println("int ", a)
	fmt.Println("int32 ", b)
	fmt.Println("int64 ", c)
}
