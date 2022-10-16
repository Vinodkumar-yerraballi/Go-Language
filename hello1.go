package main

import "fmt"

var firstname = "Vinod"
var lastname = "Kumar"
var age = 22

func main() {
	fmt.Println("%T,%v", firstname, lastname, age)

}

//another way we create variable

var (
	b = 22
	c = 44
	d = "data science"
)

func main() {
	fmt.Println(b+c, d)
}

// How to print same variable in the go

var a int = 32

func main() {
	fmt.Println(a)
	var a int = 55
	fmt.Println(a)
	var b int = 33
	fmt.Println("%v,%T", b, b)
	// use the string
	var url string = "https://paperswithcode.com/"
	fmt.Println(url)

}
