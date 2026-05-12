package main
 import "fmt"
func main(){
    // in golang is we declare a variable than we have to use it other wise we have to to delete it 
 
	//traditional methode
	var name string ="string"
	
    //infer
	//var name  ="string"
	var isAdult=true 

	fmt.Println(name)
	fmt.Println(isAdult)

	//shorthand syntax
	 language:="golang"
	 fmt.Println(language)

	 //need for the above methode to declare variable when we have shorthand syntax is 
	 // if we need to decalre a variable and after some line of code we have to assign the value to code is than we use that methode

	 var year int

	 year=2020

	 fmt.Println(year)
}