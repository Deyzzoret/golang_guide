package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

	// Functionally, both code snippets will behave the same way and produce the same output.
	// However, the first code snippet uses a more restricted scope for the randNum variable,
	// which can be beneficial if you want to limit the variable's accessibility to only the 
	// specific block where it's needed.
	// On the other hand, the second code snippet has a broader scope for randNum,
	// allowing it to be used throughout the main() function if necessary.
	
    if randNum := rand.Intn(100); randNum%2 == 0 {
        fmt.Print("Bingo!")               
    } else {
        fmt.Print(randNum)
    }       
}