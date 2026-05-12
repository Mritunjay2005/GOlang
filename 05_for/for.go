package main

import "fmt"

//for -> only construct for looping
//there is no while loop or do_while loop
func main() {
	//while lop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1

	}

	fmt.Print("\n") //this is use to skip a line

	//infinte loop **
	// for{
	//fmt.Println("1")
	// }
	//to canle from this inifnte lop we have to do "ctrl+c"

	//clasic for loop
	for i := 1; i <= 3; i++ {
		//break      -> this will break the loop

		//continue    -> this will skjip that iteration
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Print("\n")

	//1.22 update contain a coincept of range
	for i := range 3 {
		fmt.Println(i)
	}
}
