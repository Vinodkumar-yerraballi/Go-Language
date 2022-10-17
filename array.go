package main

import (
	"fmt"
)

func main() {
	marks1 := 55
	marks2 := 65
	marks3 := 94
	marks4 := 55

	fmt.Println(marks1, marks2, marks3, marks4)

	// another way define the arrays
	marks := [3]int{66, 77, 99}
	fmt.Println(marks)
	arry := [...]int{87, 80, 100, 2093, 5784, 9086}
	fmt.Println(arry)

	var student [3]string
	student[0] = " Apple"
	student[1] = "Ball"
	student[2] = "Cat"
	fmt.Println(student)

	var greeting [4]string
	greeting[0] = "Hi"
	greeting[1] = "Good morning"
	greeting[2] = "How may assist you"
	greeting[3] = "Thank you"
	fmt.Println(greeting)
	// Print the specific index given below
	fmt.Println(greeting[2])
	//checking the length
	fmt.Println(len(greeting))

	//2d dimenstion array
	var matrix2d [2][2]int
	matrix2d[0] = [2]int{0, 1}
	matrix2d[1] = [2]int{1, 0}
	fmt.Println(matrix2d)
	//create 2*3 matris
	var matrix3d [3][2]int
	matrix3d[0] = [2]int{0, 0}
	matrix3d[1] = [2]int{0, 1}
	matrix3d[2] = [2]int{1, 1}
	fmt.Println(matrix3d)
	// another way to create matrix
	var matrix4d [2][2]int = [2][2]int{[2]int{3, 5}, [2]int{6, 9}}
	fmt.Println(matrix4d)

	var matrix5d [2][3]int = [2][3]int{[3]int{1, 1, 1}, [3]int{2, 2, 2}}
	fmt.Println(matrix5d)

	arry1 := [...]int{1, 3, 4}
	arry2 := arry1
	arry2[0] = 0
	arry2[1] = 6
	fmt.Println(arry1)
	fmt.Println(arry2)
	// slices
	a1 := []int{1, 2, 4}
	fmt.Println(a1)
	fmt.Println(len(a1))
	fmt.Println("Capcity", cap(a1))
	a2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a2[:]   //[1 2 3 4 5 6 7 8 9 10]
	c := b[3:]   // [4 5 6 7 8 9 10]
	d := b[:6]   //[1 2 3 4 5 6]
	e := a2[3:6] // [4 5 6]

	// How to change the element in array
	a2[3] = 999
	fmt.Println(a2)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
