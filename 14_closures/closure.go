package main
	
import "fmt"

	// Closure is a function that can access variables from its enclosing scope,
	//  even after the outer function has finished executing.

	func counter() func() int{//if fn.() is executed then it is removed to the call stack but the value of variable is still preserved in the closure
      c:=0
	  return func()int{//ananymous function
		c++
		return c
	  }//clouser can hold the value of variables even after fn.() has finished executing
	}
func main(){

     increment:=counter()
	 fmt.Println(increment())
}