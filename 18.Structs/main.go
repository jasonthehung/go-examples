package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name, age: 42}
	return &p
}

func main() {
	fmt.Println(Person{"Bob", 20})
	
	fmt.Println(Person{name: "Alice", age: 30})

	fmt.Println(Person{name: "Fred"})

	fmt.Println(&Person{name: "Ann", age: 27})

	johnPtr := newPerson("John")
	fmt.Println("Pointer:", johnPtr)
	fmt.Println("Value:", *johnPtr)
	

	s := Person{name: "Sean", age: 50}
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