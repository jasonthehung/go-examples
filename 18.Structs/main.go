package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name, age: 42}
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	
	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 27})

	johnPtr := newPerson("John")
	fmt.Println("Pointer:", johnPtr)
	fmt.Println("Value:", *johnPtr)
	

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sPtr := &s
	fmt.Println(sPtr.age)

	sPtr.age = 51
	fmt.Println(sPtr.age)

	dog := struct {
		name string
		isGood bool
	}{
		name: "Rex",
		isGood: true,
	}
	fmt.Println(dog)
}