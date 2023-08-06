package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Declare an empty integer array
	var x [4]int

	// There are severals ways to initialize the value of 
	// an array

	var rgb [3]string = [3]string{"red", "green", "blue"}

	// Short declaration
	// rgb := [3]string{"red", "green", "blue"}

	// Let go determine the number of elements (length) of
	// array 
	// rgb := [...]string{"red", "green", "blue"}


	// Comparison of 2 arrays

	// In Go, arrays with the same element types and the same length can be compared using the == operator
	// to check if they are equal. When comparing arrays, each element in both arrays is compared individually,
	// and the arrays are considered equal if all elements are equal in both arrays and they are in the same order.

	var arr = [...]int{1, 2, 3, 4}
	var anotherArr = [4]int{1, 2, 3, 4}
	fmt.Println(arr == anotherArr)


	// In golang is not possible to use a variable to determine the size of an array


	// // Add elements to array in golang
	// In this snippet, slice is declared as an array with a fixed length of 2.
	// Since slice is an array, you cannot use the append function directly with it.
	// The append function is used to append elements to slices, not arrays.
	// Hence, the code will produce the error you mentioned: "invalid argument: slice (variable of type [2]int) is not a slice."
	// var slice = [2]int{1, 2}
	// var slice = [...]int{1, 2}

	// slice = append(slice, 2, 0, 2, 2)


	// In this snippet, slice is declared as a slice. Slices in Go are dynamic and can grow in size.
	// The append function works with slices, allowing you to add elements to the end of the slice.
	// In this case, the append function appends the elements {2, 0, 2, 2} to the slice, resulting in the slice {1, 2, 2, 0, 2, 2}.

	var slice = []int{1, 2}
	slice = append(slice, 2, 0, 2, 2)

	// Compare to slices

    z := []int {2, 0, 1, 7}
    y := []int {2, 0, 2, 2}

    fmt.Println(reflect.DeepEqual(z, y))



}