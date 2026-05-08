package main

//IMP:=this feature make go-lang a very powerfull language

import (
	"fmt"
	"time"
)

//go-rutine  is used to run program in parallel in many lightweighted threads
//which means gorutine utilizes the concept of multithreading

func task(id int) {
	fmt.Println("do task no. ", id)
}
func main() {
	for i := 0; i < 10; i++ {
		task(i)
	}
	//if we ant to run this loop in parallel than we will utilize gorutine
	fmt.Println("with gorutine")
	fmt.Printf("\n")

	//methode 1.
	for i := 0; i < 10; i++ {
		go task(i)
	} //this is printing an unorgonized data every time different
	time.Sleep(time.Second * 2)
	fmt.Printf("\n")

	//methode 2.
	for i := 0; i < 10; i++ {
		func(i int) {
			task(i)
		}(i)
	} //this also print an unorgonized data set every time different

	//since this loop will be exefcutee din milliseconds and than there is no work to do the program will shut and we will not be able to see the output so
	// we will be using a delay command
	time.Sleep(time.Second * 2)
	//the reason i used  time.Sleep(time.Second*2) two times in this program is because the process is being finished so faasts taht my loops are not devided into 10,10,10 but somthing like 10, x, 20-x-->because there is a delay into printing next empty line
}
