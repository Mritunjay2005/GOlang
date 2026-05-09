package main

//IMP:=waitGroup are used with gorutine for delay and time managment
import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {//recieve pointer
	defer w.Done()//to add buffer
	fmt.Println("do task no. ", id)
}
func main() {
	 var wg sync.WaitGroup//declaring waitgroup 

	//methode 1.
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i,&wg)//send pointer
	} 
	 wg.Wait()//to add delay
	fmt.Printf("\n")

}
