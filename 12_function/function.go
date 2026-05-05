package main

import (
	"fmt"
)

func add(a, b int) int { // standard way is to write like (a int , b int) but we can also write like (a,b int) if both variables are of same type
	//   |
	//   --> this shows that one int variable will be returned by this function
	//        if not use dthan there will be no return value

	return a + b //here we are returning the sum of a and b
}

//  multiple return

// in go we can also return multiple values from a function
// but here we are only returning one value which is the sum of a and b
func language() (string, string, int, string, bool) {
	return "golang", "is the no.", 1, " programing language. -->", true // there 9s no need yo use () for return type if we are returning only one value but if we are returning multiple values then we have to use () for return type
}

func processIt(fn func(a int) int) {
	fn(1)
}

func processIt2() func(a int) int {
	return func(a int) int {
		return 2
	}
}

func main() {

	result := add(5, 9) //here we are calling the add function and passing 5 and 9 as arguments and storing the return value in result variable
	fmt.Println(result)
	fmt.Println(language()) //here we are calling the language function and printing the return values
	// |
	// here is atre println a function that's why we are using () after it

	// if we want to store the return values of language function in separate variables then we can do like this
	a, b, c, d, e := language() // here we are storing the return values of language function in separate variables a,b,c,d,e
	fmt.Println(a, b, c, d, e)

	// if we domn't want to print a paticual value then we can use _ (underscore) to ignore that value
	_, _, c1, _, _ := language() // here we are ignoring the first, second, fourth and fifth return values of language function and storing only the third return value in c1 variable
	fmt.Println(c1)              //we use _ because if we use a variable name then it will give us an error
	// that the variable is declared but not used because we are not using that variable
	// anywhere in our code so we use _ to ignore that variable and avoid the error

	fn := func(a int) int { // here we are declaring an anonymous function and assigning it to a variable fn
		return 2
	}
	fmt.Println(fn)

	fn2 := processIt2()
	fn2(6) // here we are calling the fn2 function and passing 6 as an argument and printing the return value
	fmt.Println(fn2)

}
