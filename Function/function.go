package main

import (
	"fmt"
)

func calcAmount(price int, qty int) int {
	var totalAmount = price * qty
	return totalAmount

}

func main() {
	price, qty := 20, 5
	totalAmount := calcAmount(price, qty)
	fmt.Println(totalAmount)

}

// calculate the traingel\
func equiTraingle(base float64, height float64) (float64, float64) {
	var area = 0.5 * base * height
	var parimeter = 2 * base
	return area, parimeter

}

func main() {
	base, height := 10.5, 4.6
	area, perimeter := equiTraingle(base, height)
	fmt.Println(area, perimeter)
}
