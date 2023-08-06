package main

import (
	"fmt"
)

func main() {

	var today string
	
	// Similar to C
	fmt.Scan(&today)
	fmt.Println(today)

    switch today {
    case "Sunday":
      fmt.Println("Here we go")
    case "Monday":
      fmt.Println("We have just started")
    default:
      fmt.Println(today)
	}

	fmt.Println("Introduce your age")

	var age int
	fmt.Scan(&age)
	fmt.Println("Age's value")
	fmt.Println(age)
	fmt.Println("Age's memory value")
	fmt.Println(&age)



}
