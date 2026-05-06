package main
import "fmt"
//by value -->it mena the num does not contain the source value it contain the copy of the sourceb value 
func inChange(num int){
	num=5
	fmt.Println("the in cahnge by value is :",num)
}
//by reference
func inChange1(num *int){//by putting * before int we are showing that we are using pointer to point to the source variable
*num =5//by using *--> we are dereffernceing it
fmt.Println("in change function by reference :", *num)//if we don't use */dereffernce than it will print the address
}
func main(){
	nums:=1//source value
	inChange(nums)
    fmt.Println("in main funtion :",nums)

	//if we want to see the location in which the value is stored in memory
	//we will use & before the variable
	fmt.Println("memoru addred of nums: ",&nums)
    
	
	inChange1(&nums)
    fmt.Println("in main funtion :",nums)



}