package main

import "fmt"

// Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
func plus(x int, y int) int {
	return x + y
}

func multiPlus(x, y, z int) int {
	return x + y + z
}

func main() {
	res := plus(10, 20)
	fmt.Println("res:", res)

	res = multiPlus(60, 20, 70)
	fmt.Println("res:", res)
}