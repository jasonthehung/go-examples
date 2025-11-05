package main

import "fmt"

func sum(nums ...int) int {
	fmt.Println(nums, " ")
	total := 0

	for _, val := range nums {
		total += val
	}
	
	return total
}

func main() {
	
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3, 4, 5))

	nums := []int{6, 7, 8}
	fmt.Println(sum(nums...))
}
