package main

import (
	"fmt"
	
	"time"
)

// Structs are a way to group together related data into a single unit.
//  They are similar to classes in other programming languages,
// but they do not have methods or inheritance. Structs are defined using
// the "type" keyword, followed by the name of the struct and the fields that it contains.

//structs are used many times in a progarm

//structs embedding is a way to include one struct inside another struct.
//  This allows us to reuse the fields and methods of the embedded struct in the outer struct without having to explicitly define them in the outer struct.
// The embedded struct is also known as the "anonymous field" because it does not have a name and can be accessed directly from the outer struct.

type customer struct{
	name string
	email string
	age int
}

//example e-com order struct
type order struct{
	id string
	amout float64
	status string
	creatredAt time.Time//nanosecond precision timestamp

 //suppose we need to add customer details in the order struct then we can do like this
	customer // here we are embedding the customer struct inside the order struct which allows us to reuse the fields of the customer struct in the order struct without having to explicitly define them in the order struct
}


//methods
//  |
//   --> methods are functions that are associated with a struct and can be called on an instance of the struct.

//reciver function for order struct
func (o *order)changeStatus(status string){//here we are using poijnter
    o.status=status//we didn't use dereference operator * here because structs do this on behalf of us
}

func (o *order)getAmount() float64{
	return o.amout
}


func newOrder(id string, amount float64 , status string) *order{
	//initial setup goes here...
	myOrder:=order{
		id: id,
		amout: amount,
		status: status,
	}
	return &myOrder
}

func main(){

  user1:=customer{
	name: "John Doe",
	email: "john.doe@example.com",
	age: 30,
    }


	//if we don't want to set feild then default value is zero value of the field type
  order1 := order{
	id: "0001",
	amout: 999.99,
	status:"out for dilivery",
    customer: user1,//1st way to add
    }
    order1.creatredAt = time.Now() // here we are setting the createdAt field of the order struct to the current time using time.Now() function which returns the current time in nanosecond precision timestamp
    fmt.Println(order1)
     
    fmt.Printf("\n") // here we are printing a new line to separate the output of the two orders

   order2:=order{
	id: "0002",
	amout: 1999.99,
	status:"processing",
	creatredAt: time.Now(), // here we are setting the createdAt field of the order struct to the current time using time.Now() function which returns the current time in nanosecond precision timestamp
    customer: customer{//2nd way to add
	name: "John Doe",
	email: "john.doe@example.com",
	age: 30,
    },    
   }
   order2.customer.name="robin"// change the value 
    fmt.Println(order2)
  

	order1.changeStatus("delivered")// here we are calling the changeStatus method of the order struct to change the status of the order1 struct to delivered
	//here it will show the updated status of the order1 struct because we are using a pointer receiver for the changeStatus method which allows us to modify the original struct that is passed to the method
	fmt.Println(order1) // here we are printing the order1 struct after changing the status to delivered using the changeStatus method

    
	fmt.Println(order2.getAmount()) // here we are printing the order2 struct after changing the amount to 1500.00 using the getAmount method
    
	myOrder:=newOrder("0003", 2999.99, "processing") // here we are creating a new order using the newOrder function which returns a pointer to the order struct
	fmt.Println(myOrder) // here we are printing the myOrder struct which is a pointer to the order struct created using the newOrder function
   
   //but somethinmg we only need to use value for one time 
   //so we can directly use the struct without creating a variable for it
   language:=struct{
	name  string
	isGood bool
   }{"Golang", true} // here we are creating an anonymous struct and initializing it with values for the name and isGood fields
   fmt.Println(language) // here we are printing the language struct which is an anonymous struct created and initialized in the same line





}


//IMP: use * only when we have to modify the original struct
//     if we don't want to modify the original struct then we can directly use without *
