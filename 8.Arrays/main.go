package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	
	fmt.Println("len: ", len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl: ", b)

	c := [...]int{6,7,8,9,0}
	fmt.Println("dcl: ", c)

	// Specify the index with :, the elements in between will be zeroed.
	d := [...]int{10, 4: 99, 50}
	fmt.Println("idx: ", d)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
    	{5, 5, 5},
   		{6, 6, 6},
	}
	fmt.Println("2d: ", twoD)
}
