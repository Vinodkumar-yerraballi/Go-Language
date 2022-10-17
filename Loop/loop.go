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
	// slices
	// nums := []string{"one", "two", "three"}
	// for idx, num := range nums {
	// 	fmt.Println("%d: %s\n", idx, num)
	// }
	// for idx := range nums {
	// 	fmt.Println("%d: %s\n", idx)
	// }
	// for _, num := range nums {
	// 	fmt.Println("%d: %s\n", num)
	// }
	// arays
	// nums := [...]string{"one", "two", "three"}
	// for idx, num := range nums {
	// 	fmt.Println("%d \n", idx, num)
	// }
	// for idx := range nums {
	// 	fmt.Println("%d \n", idx)
	// }
	// for _, num := range nums {
	// 	fmt.Println("%d \n", num)
	// }
	// for range nums {
	// 	fmt.Println("Hello world")
	// }

	// map
	nums := map[string]int{"one": 1, "two": 2, "three": 3}
	for k := range nums {
		fmt.Println(k)
	}
	for k, v := range nums {
		fmt.Println("%d: %s\n", k, v)
	}

	m := []string{"Zero", "One", "Two"}
	var idx int
	for idx = range m {

	}
	fmt.Println(idx)
}
