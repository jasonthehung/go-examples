package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	x, y := vals()
	fmt.Println("x:", x, "y:", y)

	_, z := vals()
	fmt.Println("z:", z)

}
