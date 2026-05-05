package main

import (
	"fmt"
)

func main() {
	// range is used to iterate over elements in a variety of data structures.
	
	
	//slice
	nums:= []int{5,6,7,8}

	sum:=0

	for i, num:=range nums{// instyed of i we can also use _ (underscore) if we dont want to use index value
		sum+=num
		fmt.Print(i, " ")
	}
	fmt.Println(sum)


	//map
	m:=map[string]string{
       "palyer_1":"harsh",
	   "player_2":"yash",
	   "palyer_3":"krish"}

	for k,v:=range m{ //if we only use k(one variable) -->then it will store key 
		fmt.Println(v,"->",k)
	}

	//unicode code pint runes of a string
	//starting byte of rune and the rune itself
	//255< then we store in 1 byte,if not then 2 bytes
	for i,c:=range"golang"{
		fmt.Println(i,c)
	}
	fmt.Println("\n")
	
// if we want to print the character instead of unicode code point then we can use %c verb in fmt.Printf
   for j,l:=range"goalng"{
	fmt.Println(j,string(l))
   }	
}

