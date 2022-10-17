package main

import "fmt"

var k int = 4949

func main() {
	fmt.Println("Hello world!")
	fmt.Println("I'm Very happy to learn go language")
	// how to create variable
	var a int
	a = 22
	fmt.Println("The value a= ", a)
	// create the float variabel
	var b float32
	b = 3.5
	fmt.Println("The value b=", b)

	// create the string variable

	var name = "This is beging of the go language"
	fmt.Println(name)

	// another way to create variable

	c := 33
	fmt.Println(c)

	fmt.Println("c: %T\n", c)

	fmt.Println(k)
}
