package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	default:
		fmt.Println("i is neither 1 nor 2")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	whatAmI := func(i interface{}) string {
		switch t := i.(type) {
		case bool:
			return fmt.Sprintf("I'm a %T", t)
		case int:
			return fmt.Sprintf("I'm a %T", t)
		default:
			return fmt.Sprintf("Don't know type %T", t)
		}
	}
	fmt.Println(whatAmI(true))
	fmt.Println(whatAmI(42))
	fmt.Println(whatAmI("hello"))
}
