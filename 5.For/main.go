// `for` is Go’s only looping construct

package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Print(i)
		i++
	}
	fmt.Println()

	for j := 0; j < 3; j++ {
		fmt.Print(j)
	}
	fmt.Println()

	for i := range 3 {
		fmt.Print(i)
	}
	fmt.Println()

	for {
		fmt.Print("loop")
		break
	}
	fmt.Println()

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Print(n)
	}
	fmt.Println()
}
