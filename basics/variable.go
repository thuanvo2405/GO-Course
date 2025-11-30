package basics

import "fmt"

var middleName = "Minh"

// we can't use go syntax := because this is a global scope and := syntax just only use in func

func main() {
	// var age int
	// var name string = "Thuan"
	// var nickname = "Thuan Vo"

	// count := 10
	// lastName = "Vo"
	fmt.Println(middleName)

	printName()
	// DEFAULT VALUE
	// Numberic types: 0
	// Boolean types: false
	// String types: ""
	// Pointers,slices,maps,functions and structs: nil

}

func printName() {
	firstName := "Device"
	fmt.Println(firstName)
	middleName = "Test"
	fmt.Println(middleName)

}
