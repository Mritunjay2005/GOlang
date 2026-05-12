package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch case
	i := 3
	switch i {

	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")

	case 4:
		fmt.Println("four")

	default:
		fmt.Println("not found")

	}
	//switch case with multiple cases
	switch time.Now().Weekday() {

	case time.Saturday, time.Sunday: //for this time librabry we have to import the time package
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	//type switch case
	whoAmI := func(i interface{}) { // declare a functiopn that takes an empty interface as an argument
		switch t := i.(type) {
		case int:
			fmt.Println("integer")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("boolean")
		default:
			fmt.Printf("unknown type %T\n", t)
		}
	}

	//call the whoAmI function with different types of arguments
	whoAmI(42)
	whoAmI("hello")
	whoAmI(true)
	whoAmI(3.14)

}
