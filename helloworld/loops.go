package main

import (
	"fmt"
)

func main() {

	sum := 0
	for i := 0; i < 50; i++ {
		sum += i
		fmt.Println(sum)
	}

	// Infinite loop
	for {
		fmt.Print("forever!")
	}
}