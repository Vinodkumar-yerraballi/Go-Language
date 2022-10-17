package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Print("If is true\n")
	}
	// Relation operators
	a := 10
	b := 10
	if a > b {
		fmt.Println("a is greater than b")
	}
	if a < b {
		fmt.Println("b is greater than a")
	}
	if a == b {
		fmt.Println("Both are equal")
	}
	// Logical operators
	c := 9
	d := 55
	e := 99
	if c < d && d > e {
		fmt.Println("c is bigger")
	} else if d > c && e < c {
		fmt.Println("D is bigger")
	} else {
		fmt.Println("e is bigger")
	}

	if c > d || d > e {
		fmt.Println("C lies between")
	}
	if !false {
		fmt.Println("True")
	}
	if num := 100; num > 0 {
		fmt.Println("Negative number")
	}
	if num := 50; num%2 == 0 {
		fmt.Println("The number is evem")
	}

}
