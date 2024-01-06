package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mgp87/go_learning/exercises"
)

var fileName string = "./files/saved_files/file.txt"

func SaveTable() {
	var text string = exercises.MulTable()
	file, err := os.Create(fileName)
	if err != nil {
		panic("Error creating file " + err.Error())
	}
	fmt.Fprintln(file, text)
	file.Close()
}

func AddTable() {
	var text string = "\n" + exercises.MulTable()
	if !Append(fileName, text) {
		fmt.Println("Error adding text to file")
	}
}

func Append(fileName string, text string) bool {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("Error opening file " + err.Error())
	}
	_, err = file.WriteString(text)
	if err != nil {
		panic("Error writing to file " + err.Error())
	}
	file.Close()
	return true
}

func ReadFile() {
	file, err := os.Open(fileName)
	if err != nil {
		panic("Error reading file " + err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		record := scanner.Text()
		fmt.Println("> ", record)
	}
	file.Close()
}
