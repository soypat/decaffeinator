package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch today {
	case time.Saturday:
		fmt.Println("Today.")
	case time.Friday:
		fmt.Println("Tomorrow.")
	case time.Thursday:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
