package helloworld

import (
	"fmt"
	"math/rand"
	"strconv"
) 


func main() {

	// Avoid cast of interger value to ascii value instead of actual value
	// in string
    var x int = 101
    var y string
    y = strconv.Itoa(x)
    fmt.Println(y)


	// Prints a random number between 0 and 10
	// I think that in this case, we are not using the %v as we are not
	// accessing any variable's value
	fmt.Println("A random integer:", rand.Intn(10))


}