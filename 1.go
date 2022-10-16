package main

import "fmt"

func main() {
	var a bool = false
	fmt.Println("%v,%T\n", a)
	var b bool = true
	fmt.Println(b)
	n := 1 == 1
	m := 1 == 2
	fmt.Println(m)
	fmt.Println(n)

	var A int32 = 124
	fmt.Println(A)

	var B uint16 = 35
	fmt.Println(B)
	c := 133
	d := 55
	fmt.Println("%v\n", c+d)
	fmt.Println("%v\n", c-d)
	fmt.Println("%v\n", c*d)
	fmt.Println("%v\n", c%d)
	fmt.Println("%v\n", c/d)

	var e int = 66
	var f int8 = 122
	fmt.Println(e + int(f))
	var g int = 329
	var h float32 = 34.46
	fmt.Println(g + int(h))

	// bit operators

	i := 44
	j := 89
	// And operators
	fmt.Println(i & j)
	//
	fmt.Println(i ^ j)
	fmt.Println(i | j)
	fmt.Println(i &^ j)

	a := 966
	a = a << 4 // 999^4 * 999^4
	fmt.Println(a)
	a = a >> 4 // 999^4 / 999^4=999^0
	fmt.Println(a)

}
