package main

import (
	"fmt"
	"sync"
)

//Mutex and RWMutex

//These are used to communicate by sharing memory
//Bit of anti go idiom
//General Go's philosophy is to share memory by communicating

//Mutex guard critical section of programs
//Critical section : Section of code that requires exclusive access to a SHARED RESOURCE
//In following examples memory, a variable, is this critcal Section
//Developer is responsible for guarding the variable using Mutex

//In following example, there are two concurrent go-routines vying for same variable count
//One increments its value while one decrements
//We use lock to restrict access of the variable to one of the go-routine
//Once that go-routine has finished using the variable then only the other can acquire the lock and use it

//This way we have synchronised the access

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()         // acquire the lock and start guarding the data
		defer lock.Unlock() // seems absurd right? unlocking immidiately after locking,
		//but defer makes sure this line is called once the entire func is executed
		count++ // do the operation on count
		fmt.Printf("imcremeting count : %d \n", count)
	}

	decrement := func() {
		lock.Lock()         // acquire the lock and start guarding the data
		defer lock.Unlock() // seems absurd right? unlocking immidiately after locking,
		//but defer makes sure this line is called once the entire func is executed
		count-- // do the operation on count
		fmt.Printf("imcremeting count : %d \n", count)
	}

	// Launcing above defined functions as go-routines
	// adding WaitGroup so that go-routines hae chance to get executed
	var arith sync.WaitGroup

	for i := 0; i < 5; i++ {
		arith.Add(1)
		go func() {
			defer arith.Done()
			increment()
		}()
	}

	for i := 0; i < 5; i++ {
		arith.Add(1)
		go func() {
			defer arith.Done()
			decrement()
		}()
	}

	arith.Wait()

	fmt.Println("completed arithmatic")
}
