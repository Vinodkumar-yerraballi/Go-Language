package main

import (
	"fmt"
)

type Celebrity struct {
	age           int
	movie         []string
	CelebrityName string
	coActorrs     []string
}

type DataScience struct {
	Program  []string
	uses     []string
	subfield []string
}

func main() {
	acelebrity := Celebrity{
		age: 55,
		movie: []string{
			"Bahuballi",
			"Bahuballi-2",
		},
		CelebrityName: "Prabas",
		coActorrs: []string{
			"Rana",
			"Anushka",
			"Subaraju",
		},
	}
	fmt.Println(acelebrity)

	bDataScience := DataScience{
		Program: []string{"Python",
			"R",
			"C++"},
		uses: []string{
			"Health",
			"It",
			"Finace",
		},
		subfield: []string{
			"DeepLearing",
			"Machine Learing",
			"Natural Language",
		},
	}
	fmt.Println(bDataScience)
	// print the subfilds
	fmt.Println(bDataScience.subfield[1])
	// another way to create structs
	man := struct{ age int }{age: 76}
	manb := man
	manb.age = 99
	fmt.Println(man)
	fmt.Println(manb)
}
