package main
import "fmt"

func sum(nums ... int) int{// here we are declaring a variadic function sum that takes a variable number of arguments of type int and returns an int value which is the sum of the numbers passed as arguments to the function
	// here we are using ... (three dots) to indicate that the function is a variadic function and it can take a variable number of arguments of type int
	// by writeing int after ... (three dots) we are indicating that the function can take a variable number of arguments of type int and it will be treated as a slice of int inside the function
	// if we domn't want to decalre the varaibel type then we can also write like this func sum(nums ...){...} but it is not recommended because it can lead to confusion and it is not clear what type of arguments the function can take
	// so we use interface{} type to indicate that the function can take a variable number of arguments of any type but it is not recommended because it can lead to confusion and it is not clear what type of arguments the function can take
	total:=0
for _,v:= range nums{
	total+=v
}
	return total
}


// variadic function is a function that can take a variable number of arguments of the same type
// we can use ... (three dots) to indicate that a function is a variadic function	
func main(){
	
	//here println is a variadic function because it can take a n no. of variable number of arguments of different types
	fmt.Println(1,2,23,"string",true) // here we are printing multiple values of different types using fmt.Println function

    result:=sum(1,2,3,4,5) // here we are calling the sum function and passing multiple values of type int as arguments and storing the return value in result variable
	fmt.Println(result) // here we are printing the result variable which contains the sum of the numbers passed as arguments to the sum function
    

	//suppose we want to pass a slice of int to the sum function then we can do like this
	num:=[]int{5,6,7,8,9}
	result1:=sum(num...) // here we are passing the slice of int to the sum function using ... (three dots) to unpack the slice
	fmt.Println(result1) // here we are printing the result variable which contains the sum of the numbers passed as arguments to the sum function	

}