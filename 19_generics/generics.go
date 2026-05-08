package main

import (
	"fmt"
)
func printSlice(item []int){
	for _,item:=range item{
		fmt.Println(item)
	}
}
// code logic is same and we are repeating it with diffrenet datat type input ↑&↓ 
// to avoid this we can use generics which is a feature in go 1.18 and above
func printStringSlice(item []string){
	for _,item:=range item{
		fmt.Println(item)
	}
}

func printSlice1[t any](item []t){//using any is not right because it can accept any type of data
	// insted of any we can use --> interface{} but it is not recommended because it can accept any type of data and we will lose the type safety
	for _,item:=range item{
		fmt.Println(item)
	}
}
func printSlice2[t int | string ](item []t){//here we are constrening it to only two data type int and string 
		for _,item:=range item{
		fmt.Println(item)
	}
}
func printSlice3[t comparable](item []t){//here we are using comparable it contain a large varitey of datt type  
		for _,item:=range item{
		fmt.Println(item)
	}
}
//if we want to pass multiple data types in one function here 
func printSlice4[t comparable, s string, i int,](item []t , name []s, value []i){//here we are using comparable it contain a large varitey of datt type  
		for _,item:=range item{
		fmt.Println(item,name,value)
	}
}

//here is how to use generic in struct
type stack [t any]struct{
	element[]t
}
func main(){
  nums:=[]int{1,2,3}
  status:=[]bool{true,false,true}
  names:=[]string{"golang","java","python"}
  printSlice(nums)
  printStringSlice(names)
  //here we are passing both int and string
  printSlice1(nums)
  printSlice1(names)
  printSlice1(status)
  //here after we add constrant 
  printSlice2(nums)
  printSlice2(names)
  //printSlice2(status)-->it is showing error in compiler because the fn. can only thae two type of data type
 //if we want to pass boolean data type also then we need to add bool also in that generic

 //here we are passing throug comparable type
  printSlice3(nums)
  printSlice3(names)
  printSlice3(status)
//here we are passing multiple data types in one function here
printSlice4(status,names,nums)

 //struct usage
 mystack:=stack[string]{//here we have to define what dta type is stack 
	element: []string{"golang"},
 }
fmt.Println(mystack)


}