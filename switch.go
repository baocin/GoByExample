// switch
//This is the good commit
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Congrats!~ It's the weekend!")
	default:
		fmt.Println("Another boring weekday")
	}

	t := time.Now()
	//alternative if statement
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon!")
	default:
		fmt.Println("It's after noon...")
	}
}
