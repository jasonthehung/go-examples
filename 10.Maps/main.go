package main

import (
	"fmt"
	"maps"
)

func main() {
	// m := map[string]int{}
	m := make(map[string]int)


	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map.
	val, prs := m["k2"]
	fmt.Println("val:", val, "prs:", prs)

	m1 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", m1)

	m2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(m1, m2) {
		fmt.Println("m1 is equals to m2")
	}
}