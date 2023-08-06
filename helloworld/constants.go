package main

func main() {

	// In case we are assigning some value to a constant from
	// variables, it will raise an error as constant's value will
	// depend on variables' result

	// Constants cannot be declared using :=

	// Single constant
	const x int = 2732

	// Multiple constants in one block
	const (
		y int = 2017
		z int = 2022
	)
}