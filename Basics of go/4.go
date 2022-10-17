package main

import (
	"fmt"
)

const (
	isPrincipal = 1 << iota
	isCollege
	canGOAccount

	canGoCSE
	canGoISE
	canGoMBA
	canGoESE
)

func main() {
	var roles byte = isPrincipal | canGOAccount | canGoMBA
	fmt.Println(roles)
	fmt.Println(isPrincipal&roles == isPrincipal)
	fmt.Println(canGOAccount&roles == canGOAccount)

}
