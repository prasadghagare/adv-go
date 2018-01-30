package main
import (
  "fmt"
  "sync"
  "time"
)

func main(){
  c := sync.NewCond(&sync.Mutex{})
  //NewCond is useful when we need a go-routine to wait on some result
  //NewCond takes a type which satisfies Locker interface(this has methods Lock and Unlock)
  //Usually, Mutex and RWMutex are the types which satisfy Locker interface(that is they implement Lock/Unlock methods)

  //Cond is a type returned by NewCond with a Locker value

  //So on line 8, we got a Cond type varaiable c with Locker interface value provided by &sync.Mutex{}

  queue := make([]interface{}, 0, 10)
  //here we initialized a queue with capacity 10 but no inital values that is length is zero

  removeFromQueue := func(delay time.Duration) {

    //this routine will wait for designated time and deque the queue
    //also it will send out the signal that something has happened
    time.Sleep(delay)
    c.L.Lock()
    queue = queue[1:]
    fmt.Println("Removed from queue")
    c.L.Unlock()
    c.Signal()
}

for i := 0; i < 10; i++{
    c.L.Lock()
    for len(queue) == 3 {
	fmt.Println("lets wait")
        c.Wait()
    }
    fmt.Println("Adding to queue")
    queue = append(queue, struct{}{})
    go removeFromQueue(5*time.Second)
    c.L.Unlock()
}

//analysis:
//as long as there are not 3 elements in queue. append(queue) is called
//once len(queue) == 3 is true, we call c.Wait
//This suspends the current go-routine that is main go-routine untill it will recieve a Signal
//now by this time we have already scheduled removeFromQueue go-routine to run 3 times
//Suspending main go-routine will bring one of these in execution
//this one go-routine will deque the queue and sends signal to main go-routine which is waiting for this signal
//upon recieving this signal, we resume from the for loop
// again one more item will be added to queue and one more removeFromQueue routine wil \l be scheduled to run
//next iteration of outer for will begin
//again length of queue will be 3
// repeat from line 47
//this will continue till outer for loop completes
//last 2 removeFromQueue will never come into execution as main routine will exit

}
