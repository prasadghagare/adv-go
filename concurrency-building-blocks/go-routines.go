package main

import "fmt"
import "time"

func main() {
	//go-routines are functions that are executing concurrently with other code.
	//various ways to start go-routines
	go sayHello() // a normal function as defined below called with go keyword

	//Following construct is called as anonymous function
	//Look at the () on line 14 we are valling the function at the same time we have defined it.
	go func() {
		fmt.Println("baee- bye")
	}()

	// third way
	// define the function without the name and assign it to a variable
	sayHola := func() {
		fmt.Println("hello in spanish")
	}

	go sayHola()

	//adding time sleep so that go-routines have chance to execute
	time.Sleep(1000 * time.Millisecond)
}

func sayHello() {
	fmt.Println("hello")
}

//co-routines - piece of code that can stop execution waiting for input to arrive(i.e. they are blocked) and resume when input is available.
//co-routines are non-preemptive(something that cannot be interrupted(if they are not themselves blocked)).
//go-routines do NOT define their own suspension or resumption
//go runtime observes the behaviour of go-routines and suspends them if they are blocked or resumes when they are unblocked

//green threads - threads managed by langunage runtime
//go-routines(100) --> go run time(this has green threads)(80 green threads) --> OS threads(70)
//           P                                                       M         -->     N

//M:N Schedular M green threads are mapped to N OS threads
//go-routines are scheduled to run on Green threads
//when go-routines > green threads, scheduler distributes the go-routines across the green threads
//when some go-routines blocks others are scheduled to run.

//fork-join model of concurrency
//when a go-routine is created it splits off as a child of parent go-routine(fork)
//The goroutine will be created and scheduled with Go’s runtime to execute
//it is now expected to run concurrently with the parent go-routine on different green thread
//meanwhile, the parent's code is still executing
//At some future indeterminate point the child go-routine is expected to get executed
//If the child go-routine rejoins the parent go-routine it is called as join(join)

//If there is no join point, there is no guarantee that the child goroutine is executed.
//As soon as main go-routine finishes the program execution ends.
//In above example, child go-routines were able to execute because there was sleep introduced.
//But sleep is not acutal join , it introduces race condition.

//see next example to see how can we create join in go programs.
