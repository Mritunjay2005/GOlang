package main

import "fmt"

func main() {

	age := 80
	// if else working flow is if{true}-->staement1
	//                            |
	//                         {fasle}-->else if{true}-->statement2
	//                                             |
	//                                          {false}-->else --> statement3
	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior.")
	}
	fmt.Print("\n")
	fmt.Println("does person have permission to access the system?")

	var role = "admin"
	var hasPermissions = true

	if role == "admin" || hasPermissions {
		fmt.Println("Yes")
	}
	if role == "admin" && hasPermissions {
		fmt.Println("Yes root user")
	}
	//  &&--> is for and operator  --->  it merans both condition should be true to execute the statement
	//  ||--> is for or operator  ---> it means if any one of the condition is true then it will execute the statement
	//   "!"--> this is use to negation --> it is use to reverse the value of a boolean variable

	//we can declare variable in the if statement also & we can also use that varaivble in the condition of if statement
	if gender := "male"; gender == "male" {
		fmt.Println("You are a ", gender)
	} else if gender == "female" {
		fmt.Println("You are a ", gender)
	}
	//GO does not support ternary operator (condition ? true : false) like in other programming languages,
	// but we can achieve similar functionality using if-else statements or by using a function that returns a value based on a condition.
	//till now in version 1.22 there is no support for ternary operator in GO

}
