package arrays_slices

import "fmt"

var table [10]int
var table2 [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
var matrix [20][30]int
var tableSlice []int = []int{1, 2, 3}

func ShowArrays() {
	table[7] = 7
	table[2] = 2

	table3 := [10]string{"a", "b", "c", "d", "e"}

	matrix[7][24] = 17

	fmt.Println(table)
	fmt.Println(table2)
	fmt.Println(table3)
	fmt.Println(matrix)

	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}
}

func ShowSlice() {
	fmt.Println(tableSlice)

	piece := table2[3:]   // slice of table from index 3 to the end
	piece2 := table2[:5]  // slice of table from index 0 to 5
	piece3 := table2[6:8] // slice of table from index 6 to 8

	fmt.Println(piece)
	fmt.Println(piece2)
	fmt.Println(piece3)
}

func CapacityAndLength() {
	elems := make([]int, 5, 20) // slice with 5 elements and capacity of 20 (allocated space)
	fmt.Printf("Length: %d, Capacity: %d \n", len(elems), cap(elems))

	nums := make([]int, 0) // Dynamic length and capacity slice, it's like "make([]int, 0)"
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
		fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Len 100, Cap 128
	}
}
