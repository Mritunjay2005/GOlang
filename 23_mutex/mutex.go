package main

// A mutex (mutual exclusion) is a lock that ensures only one goroutine can access a piece of code or data at a time.
// Think of it like a bathroom key:
// If one person has the key → others must wait
// When they’re done → they return the key → next person can use it
import (
	"fmt"
	"sync"
)

type post struct {
	view int
	mu   sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock() //good practice
		wg.Done()
	}()
	p.mu.Lock() //Don’t lock more code than needed:
	p.view += 1
	//p.mu.Unlock() it is not a good practice
	//Mutex solves safety, but:
	// Too many locks → slow performance
	// Bad design → complex bugs
}
func main() {
	myPost := post{view: 0}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println(myPost.view)
	//IMP:-Go gives you two main ways to handle concurrency:

	// 1. Shared memory + Mutex
	// Multiple goroutines share variables
	// Use mutex to protect access

	// 👉 “Do not communicate by sharing memory; share memory by communicating”

	// 2. Channels (preferred Go style)

	// Instead of sharing data:

	// Pass data between goroutines using channels
	// Avoids many mutex problems

	// Mutex vs Channels (when to use what)

	// Situation	                      Use
	// Shared counter / state	          Mutex
	// Passing data between goroutines	  Channels
	// Complex coordination	              Channels
	// Simple critical section	          Mutex
}
