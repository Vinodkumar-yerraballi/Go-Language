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
	n := 4359.6
	n = 3.45e75
	n = 95.98E75
	fmt.Println(n)

	b := 45.45
	c := 435.2
	// Remainder used only for interges
	fmt.Println(b + c)
	fmt.Println(b - c)
	fmt.Println(b / c)
	fmt.Println(b * c)

	// Complexity

	var f complex64 = 1 + 2i
	fmt.Println(f)
	var d complex64 = 1 + 3i
	var e complex64 = 1 + 6i
	fmt.Println(d + e)
	fmt.Println(d - e)
	fmt.Println(d * e)
	fmt.Println(d / e)

}
