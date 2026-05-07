package main

import (
	"fmt"
)

//interface is a type that defines a set of method signatures.
//  A type that implements all the methods of an interface is said to satisfy that interface.
//  Interfaces are used to achieve polymorphism in Go, allowing different types to be treated as the same type through the use of interfaces.

// method 3
type paymenter interface {
	pay(amt float64) // it work like a contract
}
type payment struct {
	//methode 2
	//gatway strip//this is a concreate implementation of paymenter interface but it is not flexible because if we want to switch to razorpay then we have to change the code here also so we do somthing like ↴
	//method 3
	gatway paymenter // this is a interface type and we can assign any struct that implements the paymenter interface to this variable
}

func (p payment) makePayement(amt float64) {

	//razorpayPaymentGw :=razorpay{}//craeting instance of razorpay struct
	//razorpayPaymentGw.pay(amt)// calling the pay method of razorpay struct to make payment

	//methode 1
	//suppose the company wants to switch to strip ↴
	//stripPaymentGw:=strip{}//creating instance of strip struct
	//stripPaymentGw.pay(amt)// calling the pay method of strip struct to make payment

	//doing this way we can switch btween payment gateway but
	//it it is not flexible and we have to change the code every time we want to switch the payment gateway
	//so we do somthing lik ↴
	// methode 2↴
	p.gatway.pay(amt)
}

type razorpay struct{}

func (r razorpay) pay(amt float64) {
	// logic to make payment using razorpay
	fmt.Printf("Payment of %f made using Razorpay\n", amt)
}

type strip struct{}

func (s strip) pay(amt float64) {
	// logic to make payment using strip
	fmt.Printf("Payment of %f made using Strip\n", amt)
}

type paypal struct{}

func (p paypal) pay(amt float64) {
	// logic to make payment using paypal
	fmt.Printf("Payment of %f made using Paypal\n", amt)
}
func main() {
	//stripPaymentGw := strip{}
	// now we want to switch to paypal then we can do like this ↴
	paypalPayementGw := paypal{}
	pay := payment{
		//gatway: stripPaymentGw,
		gatway: paypalPayementGw,
	}

	pay.makePayement(100.0)
}

// here waht happened was that when we create a interface and then we create a struct that implements that interface then we can use that struct to call the method
//here we are not using implement like w e do in other programming language because
//in go it understand that ifa struct/method have same name and parameter then it is implementing that interface
