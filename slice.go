package main

import (
	"fmt"
)

func main() {
	s := make([]int, 3)
	fmt.Println(s)
	fmt.Println("Lenght", len(s))
	fmt.Println("Capacity", cap(s))

	a := []int{}
	fmt.Println(s)
	fmt.Println("Lenght", len(a))
	fmt.Println("Capacity", cap(a))
	a = append(a, 5)
	fmt.Println(s)
	fmt.Println("Lenght", len(a))
	fmt.Println("Capacity", cap(a))
	a = append(a, 3, 57, 9, 83, 56, 7, 4)
	fmt.Println(s)
	fmt.Println("Lenght", len(a))
	fmt.Println("Capacity", cap(a))

	b := make([]int, 5, 100)
	fmt.Println(b)
	fmt.Println("Lenght", len(b))
	fmt.Println("Capacity", cap(b))
	// append the elements
	b = append(b, []int{3, 5, 6, 7, 7, 8}...)
	fmt.Println(b)
	fmt.Println("Lenght", len(b))
	fmt.Println("Capacity", cap(b))
	// How to remove the element in slices from beginig
	c := []int{1, 2, 3, 4, 5}
	d := c[1:]
	fmt.Println(d)
	//Remove the last element
	e := c[:len(c)-1]
	fmt.Println(e)

	// [2 3 4 5]
	// [1 2 3 4]
	// [1 2 1 2 3]
	// this is way to use append function
	a1 := append(c[:2], c[:3]...)
	fmt.Println(a1)
}
