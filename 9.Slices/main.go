package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))


	// Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := c[2:5]
	fmt.Println("sl1:", l)

	l = c[:4]
	fmt.Println("sl2:", l)

	l = c[2:]
	fmt.Println("sl3:", l)

	t := []string{"a", "n", "n"}
	q := []string{"a", "n", "n"}
	if slices.Equal(t, q) {
		fmt.Println("t == q")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
}
