package helloworld

import "fmt"

// Package level variable
// This kind of variables can be used in all the functions in the same file,
// or it can be exported from other files
var userName string = "UserTest"

func main() {

	// It will be the same if we had declared the variable as follows:
	// var x = 4
	var x int = 4
	// Shor variable declaration is only possible inside functions
	// Golang will determine the value's type
	y := 5

	// I would say that the main difference between the Println and Printf functions,
	// is that the first one add \n add the end of each printed line. 
	fmt.Println(y)
	fmt.Println(x)


	// Declare more than one variable using only one var statement 
	
	var (
		x1 bool   = false
		y1 int    = 0
		z1 string = "false"
	)
	
	// Like C
	fmt.Printf("The type of x: %T. The value of x: %v\n", x1, x1)
}