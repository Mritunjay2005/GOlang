package main

import (
	"fmt"
	"slices"
)

func main() {

//clice->dynamic array
//slice is a reference type
//slice is a wrapper around an array
//slice is a view of an array
//slice is a pointer to an array
//slice is a descriptor of an array

//mostly used in go
//uninitiazed slice is nil
var nums []int
fmt.Println(nums) //[]
fmt.Println(nums == nil) //true
fmt.Println(len(nums)) //0

var num =make([]int, 5) //make a slice of int with length 5 -->this initialization will create an array of 5 int and return a slice that points to that array
fmt.Println(num)

//cap-->capacity of the slice-->the number of elements that the slice can hold without allocating a new array
fmt.Println(cap(num)) //5
fmt.Println(num==nil) //false

//append-->add elements to the slice
num = append(num, 6) //append 6 to the slice num
fmt.Println(num) //[0 0 0 0 0 6]
fmt.Println(cap(num)) //10-->when we append an element to a slice and the capacity is not enough to hold the new element,
//  go will create a new array with double the capacity and copy the old elements to the new array and return a new slice that points to the new array
//so when we append an element to a slice and the capacity is not enough to hold the new element, go will create a new array with double the capacity and copy the old elements to the new array and return a new slice that points to the new array
//when the capax+city is reached the capacity then the capacity will be doubled and the old elements will be copied to the new array and the new element will be added to the new array and a new slice will be returned that points to the new array




var nums2 =make([]int, 0, 5) //make a slice of int with length 0 and capacity 5
fmt.Println(nums2)



num3 := []int{} //another way to declarre a slice of int with length 0 and capacity 0
num3 = append(num3,2)//2 iws is aded
fmt.Println(num3)//[2]
fmt.Println(cap(num3))//1
fmt.Println(len(num3))//1


var nums4 =make([]int, 5, 10) //make a slice of int with length 5 and capacity 10
nums4=append(nums4,5)//append 6 to the slice nums4
var nums5 =make([]int, len(nums4), cap(nums4)) //make a slice of int with length 5 and capacity 10
fmt.Println(nums5)

//copy-->copy elements from one slice to another slice
copy(nums5, nums4) //copy elements from nums4 to nums5
fmt.Println(nums4)
fmt.Println(nums5)//nums4 data is copied to nums5

//slice operator
//slice operator is used to create a new slice from an existing slice

var nums6=[]int{1,2,3}

fmt.Println(nums6[0:2])//[1 2]-->slice operator is used to create a new slice from an existing slice
//this will exclude the last elemnet which u mentioned here
//{1. 2 . 3 }
// 0  1  2 -->position of the elements in the array
// [0:2]--> will print 0,1 and exclude 2
//so when we used nums6[0:2] it will create a new slice that points to the same array as nums6 but with a different length and capacity
//the new slice will have a length of 2 and a capacity of 3

fmt.Println(nums6[:2])//[1 2]-->this is another way to use the slice operator
//when we used nums6[:2] it will create a new slice that points to the same array as nums6 but with a different length and capacity
// beacause we did not specify the starting index, it will default to 0 and it have ending index of 2

fmt.Println(nums6[1:])//[2 3]-->this is another way to use the slice operator
//when we used nums6[1:] it will create a new slice that points to the same array as nums6 but with a different length and capacity
// because we did not specify the ending index, it will default to the length of the slice and it have starting index of 1


fmt.Println(slices.Equal(nums4, nums5))//true-->this will check if the elements of nums4 and nums5 are equal or not
// we have to import thus sclices package to use this function


//2-d slice
var nums7=[][]int{{1,2,3},{4,5,6}}//slice of slice of int
fmt.Println(nums7)
}