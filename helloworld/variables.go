package main

import "fmt"

func main() {
	// Define a variable with the string value of "Hello, World"

	var some_string string = "Hello, World"
	var hello string = "There"

	some_string2 := "Hello, World"

	// Equals to
	// fmt.Println("Hello %v", hello)

	fmt.Println("Hello", hello)

	// Print the length of a character in Go
	fmt.Println(len("Hello" + hello))

	// As in python , we can print a given number of characters by slicing the 
	// string

	fmt.Println(hello[3:])


}