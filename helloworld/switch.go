package main

import (
	"fmt"
	"time"
)

func main() {

    today := time.Now().Weekday().String()

    switch today {
    case "Sunday":
      fmt.Println("Here we go")
    case "Monday":
      fmt.Println("We have just started")
    default:
      fmt.Println(today)
	}
}