package main

import (
	"fmt"
)

func main() {
	student1 := 80
	student2 := 65
	student3 := 88
	if student1 < student2 {
		fmt.Println("Student1 get high marks compare to the student2")
	} else if student2 > student3 {
		fmt.Println("The studnet 2  get high marks student3")

	} else {
		fmt.Print("The student 3 is high marks")
	}
	var a int
	fmt.Scanln(&a)
	fmt.Println("The number is", a)
	var b int
	fmt.Println("Please enter your marks")
	if b > 35 {
		fmt.Scanln(&b)
		fmt.Println("The person is failed..!")
	} else {
		fmt.Scanln(&b)
		fmt.Println("The person is passed,Let's enjoye")
	}

	fmt.Println("Enter your age:")
	var age int
	if age > 60 {
		fmt.Scanln(&age)
		fmt.Println("Middle aged person")
	} else {
		fmt.Scanln(&age)
		fmt.Println("Senior citizen")
	}

	fmt.Println("Enter your name:")
	var name string
	var age1 int
	fmt.Scanln(&name)
	fmt.Scanln(age1)
	if name == "Vinod" {
		fmt.Scanln(&name)
		fmt.Println("your name", name)
	}
	if age1 > 30 {
		fmt.Scanln(&age1)
		fmt.Println("You are young aged person")
	} else {
		fmt.Println("You are senior aged person")
	}
}
