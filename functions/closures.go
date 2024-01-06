package functions

func table(num int) func() int {
	number := num
	seq := 0            // Initializes only the first time it's called and keeps the value stored
	return func() int { // Executes every time it's called
		seq++
		return number * seq
	}
}

func CallClosure() {
	tableOfTwo := 2
	MyTable := table(tableOfTwo)

	for i := 0; i <= 10; i++ {
		println(MyTable())
	}
}
