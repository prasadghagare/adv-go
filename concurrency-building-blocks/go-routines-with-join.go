package main

import "fmt"
import "sync"

func main() {

	//a varaible wg to tell this parent go-routine to wait for the child go-routine to finishes
	var wg sync.WaitGroup

	sayHola := func() {
		//adding relevent code on child side for stopping parent
		defer wg.Done()
		fmt.Println("Hi IN Spanish")
	}

	wg.Add(1)
	go sayHola()

	// go-routines executes in the same address space they are created in
	//what this means :) see example below

	//my_name - a variable in this address space
	my_name := "prasad"

	//Telling the parent go-routine that one more child goroutine is being added.
	wg.Add(1)
	//we create a go-routine in the same address space as current go-routine
	go func() {
		defer wg.Done()
		my_name = "Jon"
	}()

	fmt.Println(my_name)
	wg.Wait() // this is the join point - it means all go-routines are finished then only code after this point is execued.
	// thus Jon will always be printed last , because of following line
	fmt.Println(my_name)

	//starting new concept example : How go runtime holds the mmory reference required by go-routines

	//In following ex: for loop is done executing before any of the sub-gourotines it creates
	//So  ideally the reference to salutation should not be there
	//But go runtime is observent enough to hold reference to salutation variable
	// but it will be the last value that gets assigned to the variable : in this case "good day"
	//Thus when go-routines actually start the execution the value for saluation they have is "good day"
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}

	//again we wait for go routines created by for loops to end.
	wg.Wait()

	//If we want each go-routine to have one of the value of the string struct
	//we passed the salutaion one by one to each sub go-routines
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salute string) {
			defer wg.Done()
			fmt.Println(salute)
		}(salutation) //passed saluation created in for loop here
	}

	//Again wait for this for loop to end
	wg.Wait()
}
