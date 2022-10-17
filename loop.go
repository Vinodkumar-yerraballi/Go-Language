package main

import (
	"fmt"
)

func main() {
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += 1
	// }
	// fmt.Println(sum)
	// a := 40
	// i := 5
	// for i < 10 {
	// 	sum += 1
	// 	i++
	// }
	// fmt.Println(a)
	// // while loop

	// b := 4
	// for b < 100 {
	// 	b += b

	// }
	// fmt.Println(b)
	// break statmet in loop
	c := 1
	for {
		c += c
		if c == 8 {
			continue
		}
		fmt.Println(c)
		if c == 32 {
			break
		}
	}
}
