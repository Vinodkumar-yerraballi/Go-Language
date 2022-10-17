package main

import (
	"fmt"
	"time"
)

func main() {
	a := 2

	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("defalut")
	}

	i := 5
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("defalut")
	}

	// check the data is weekday or not
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekday")
	default:
		fmt.Println("This not weekday")
		break
		fmt.Println("Hello world")

	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	case t.Hour() == 12:
		fallthrough
	default:
		fmt.Println("It's after noon")

	}
}
