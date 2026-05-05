package main 
import (
	"fmt"
	"maps"
)

//map--> hash table in java, c, c++, dictionary in python ,object in js, associative array in php
func main(){
//creating map 

//1. using make function
m :=make(map[string]string)


//setting an elemnet in map
m["name"]="golang"
m["area"]="backend"


//getting an element from map
fmt.Println(m["name"], m["area"])

//IMP: if key does not exist in map then it will return zero value of that type
//for diffrent data typle string--> empty string, int-->0, bool-->false
fmt.Println(m["age"])

n:=make(map[string]int)
n["phone"]=1234567890
n["price"]=1000
n["age"]=5
fmt.Println(n["phone"],n["age"])
fmt.Println(len(n))
 delete(n,"age") //deleting an element from map
 fmt.Println(len(n))//the count will also decrease after deleting an element from map
 fmt.Println(n["age"]) //after deleting the key it will return zero value of that type

 clear(n)//to clear all the elements from map
 fmt.Println(n)//after clearing the map it will return empty map


 //2. using map literal
 m1 := map[string]string{
	 "name":"golang",
	 "area":"backend",
 }
 fmt.Println(m1)

 v,ok := m1["name"] //to check if key exist in map or not
 fmt.Println(v) //return value of key
 if ok{//ok will conatin boolean value true if key exist in map otherwise false
	fmt.Println("ok") //true
 } else{
	fmt.Println("not ok") //false
 }

 //check equal or not with maps package 
 fmt.Println(maps.Equal(m,m1)) //true
// thsi check the key and vales both are equal or not if both are equal then it will return true otherwise false
}
