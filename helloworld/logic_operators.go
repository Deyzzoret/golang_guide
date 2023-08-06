package main

import "fmt"

func main() {
    x := 2017

    fmt.Println(x > 2022 || x < 3000)
    fmt.Println(x > 2000 || x < 3000)
    fmt.Println(x > 2000 || x&2 == 0) 
	
	
	y := true
        
    fmt.Println(y)
    fmt.Println(!y)
    fmt.Println(!y && y)
    fmt.Println(!y || y)   
}