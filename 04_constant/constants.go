package main
import "fmt"
//can be declare outside of the function 
const age=30
//but using shorthand syntax   :=
//we can not decalre the const but withn trraditional methode we can declare 
var  language string="golang"
func main(){
	const name ="golang"
    // const is use when we don't wnat to change the value of the variable in the program 
	//name ="javscript"

	//const age =30

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(language)


	//group declaration of const
	const(
		port =5000
		host ="localhost"
	)
	fmt.Println(port,host)
}