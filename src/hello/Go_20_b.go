package main

import (
	"fmt"
	"time"
	"sync"
)

/*
Lesson 20_b : Goroutine Synchronization with Concurrency :

In Go : Concurrency is implemented using GO routine

NOTE NOTE :
Defer in "GO" ..  if .. waitGroup.Done() .. is not executed .. then it will enter DEADLOCK situation
 ---- Try this by commenting // waitGroup.Done() : It will result in DEADLOCK situation


 */

 var waitGroup sync.WaitGroup // WaitGroup, used to notify .. when the go routine is finished

func main()  {

	// Both the thread/GoRoutine are started concurrently .. [Not linear] & when one is completed, Done is notified to WaitGroup
	// & Then another GoRoutine will go & execute the method

	waitGroup.Add(1)
 	go say("Hey")  // Added Go to make .. Test by removing "go"

 	waitGroup.Add(1)
 	go say("There")

	waitGroup.Wait()

	/*
	Above from this code, will be completed executed concurrenlty & then below code will be executed
	 */
	waitGroup.Add(1)
	go say("Hi")
	waitGroup.Wait()
}

func say(s string)  {
	defer waitGroup.Done()  // DEFER is added, So at any cost this statement is executed, so we will make is DEADLOCK free even
	for i := 0 ; i <3 ; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}

}