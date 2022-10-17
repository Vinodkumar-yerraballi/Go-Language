package main

import (
	"fmt"
)

func main() {
	const a int = 32
	const b float64 = 4.5
	const c string = "Hello world"
	const d bool = false

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	// we can't add teo different tyes of element we get an error
	// a1 + a2 (mismatched types int and int16)
	//const a1 int = 2846
	//const a2 int16 = 23874
	const a1 int = 2846
	const a2 int = 23874
	fmt.Println(a1 + a2)
	// we can add constant variable and normal variabels

	const e int = 88
	var f int = 2829
	fmt.Println(e + f)
	// We can add constant variable
	const b1 = 33
	var b2 int16 = 444
	fmt.Println(b1 + b2)

	const c1 = iota
	fmt.Println(c1)
	const (
		d1 = iota
		e1 = iota
	)
	fmt.Println(d1, e1)

	const (
		mbaHod = iota
		cseHod
		iseHod
	)

	var HOD int
	fmt.Println(HOD == mbaHod)
	fmt.Println(HOD == cseHod)

	const (
		vinod = 0
		kumar = 1
		_     = iota
		KB    = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	var name int = 1
	fileSize := 4000000000.
	fmt.Println(name == 1)
	fmt.Println("%.2fGB\n", fileSize/GB)
}
