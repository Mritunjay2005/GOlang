package main

// channels are the pipline which are used  for communication between  go-rutine
import (
	"fmt"
	//"math/rand"
	//"time"
)

// func proccessNum(numChan chan int){
// 	//fmt.Println("number passed is ", <-numChan)
//    for num:=range numChan{
// 	fmt.Println("numbers are ",num)
//     //in range we cond't need to use '<-' to recieve message through channel
//    }
// }

// func sum(result chan int,n1 int ,n2 int){
// 	numResult:=n1+n2
// 	result<-numResult
// }

// func process(done chan bool){
// 	defer func(){done<-true}()
// 	fmt.Println("processing...")
// }
//func emailSender(emailChan <-chan string,done chan<- bool)
//  here waht we have done is we decalre that emailChan will only recerive inputs &
//  done will only send output

// func emailSender(emailChan chan string,done chan bool)
//our worker will print email from the queue one by one
// 	for email:=range emailChan{
// 		fmt.Println("sending email to ",email)
// 		time.Sleep(time.Second)
// 	}
// 	defer func(){done<-true}()
// }

func main() {
	//multi value channel
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "golang"
	}()
	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("value from channel 1 is ", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("value from channel 2 is ", chan2Val)
		}
	}

	//buffered channel
	// emailChan:=make(chan string,100)
	// done:=make(chan bool)
	//  go emailSender(emailChan,done)//trigger the func
	//  for i:=0; i<=10;i++{
	// 	emailChan<-fmt.Sprintf("%d@gmail.com",i)
	//  }

	//  fmt.Println("done sending...")//this will indicate that all the emails have been send to queue
	//  close(emailChan)//we need to close the channel otherwise their will be deadlock-> because the func will run ann infinite loop
	//  <-done

	//un-buffered channel

	//  done:=make(chan bool)
	//   go process(done)
	//   <-done//blocking //here we are juct receving
	// we use this when we want to complete the execution of oa function
	//it us somthing like waitgroup or time.Sleep(time.second)

	//  result :=make(chan int)
	//  go sum(result,5,6)
	//  fmt.Println("sum result",<-result)

	//   numChan:=make(chan int )
	//    go proccessNum(numChan)
	//    //numChan <- 2005
	//    //time.Sleep(time.Second*2)

	//    for{
	//         numChan<-rand.Intn(100)//it will generate and send random integer number from the rang of 0 to 100
	//    }

	//IMP:- witg+hout go routine using channel it will create a "deadlock"

	// messageChannnel :=make(chan string) //declaring channel

	// messageChannnel <- "golang"//blocking // send message to channel

	// msg := <- messageChannnel //  receuve message through channel
	// fmt.Println(msg)
}
