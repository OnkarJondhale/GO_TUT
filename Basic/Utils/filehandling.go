package Utils

import (
	"fmt"
	"os"
)


func Filehandling() {

	// create
	file,err := os.Create("file.txt")
	if err!=nil {
		fmt.Println("Error in file creation")
		return;
	} else {
		defer file.Close()
		fmt.Println("File creation successful")
	}

	// write
	_,_ = file.WriteString("Hello, World!")

	// read 
	content,_ := os.ReadFile("file.txt");
	fmt.Println(content)
	fmt.Println(string(content))
}