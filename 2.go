package main

import (
	"fmt"
)

func main() {

	// Real and imaginary parts
	var a complex64 = 1 + 5i
	fmt.Println(real(a), imag(a))

	var b complex64 = complex(4, 55)
	fmt.Println("%v,%T\n", b, b)
	//  string are immutable mostly not change
	s := "Hello world"
	fmt.Println(s)
	// in indexing given the values ex.108
	fmt.Println(s[2])
	// index the  leters is given  ex.l
	fmt.Println(string(s[2]))

	// we can add two string
	s1 := "Good marning"
	s2 := " How can i help you"
	fmt.Println(s1 + s2)

	// convert the string into bytes we get the husky values
	// the values of examples [71 111 111 100 32 109 97 114 110 105 110 103]
	c := []byte(s1)
	fmt.Println(c)

	// Note when we create string use double eniwited commas
	// roune is create in sigle invited commas

	r := 'r'
	fmt.Println(r)

	var k rune = 'k'
	fmt.Println(k)

	// create constant
	// once you create consta we can't change it
	const FirstName string = "Vinod"
	fmt.Println(FirstName)
}
