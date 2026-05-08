package main

import (
	"fmt"
)
//enumerated type

//this is used fo integer type only if orderStatus is of int type
type orderStatus int//here we are delaring a data type 
const(//here we are assigining to data type
	prepared orderStatus = iota
	inShipping
	outForDelivery
	delivered
	received// we can also write received orderStatus = iota but it is not necessary as it will automatically take the value of 4 as it is the next value after delivered which is 3
)
//this is how we use string for enumerated type if orderStatus is of string type
type orderStatusString string
const(
	preparedS orderStatusString = "prepared"
	inShippingS  = "inShipping"
	outForDeliveryS  = "outForDelivery"
	deliveredS  = "delivered"
	receivedS  = "received" // we can also write receivedS orderStatusString = "received" but it is not necessary as it will automatically take the value of "received" as it is the next value after deliveredS which is "delivered"
)
//calling fn. for both types are same but the type of argument is different
// in main and changeOrderStatus fn.

func changeOrderStatus(status orderStatus){
	fmt.Println("order status changed to ",status)
}
func changeOrderStatus1(status orderStatusString){
	fmt.Println("order status changed to ",status)
}
func main(){
  changeOrderStatus(received)
  changeOrderStatus1(receivedS)
}