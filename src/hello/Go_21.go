package main

import (
	"fmt"
	"time"
	"sync"
)

/*
Lesson 21 : PANIC & RECOVER

Panic : Hold available to programmer on certain condition : Program can stop running
recover(): In case of Panic, it is smoothy handled in recover(r) -- r := recover() .. if this is NOT null ..
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
	defer cleanup()
	for i := 0 ; i <3 ; i++ {
		time.Sleep(time.Millisecond*100)
		fmt.Println(s)
		if i == 2 {
			panic("oh dear, it is  2")
		}
	}
}

func cleanup()  {
	if r := recover();	r != nil {
		fmt.Println("Recovered in cleanup :", r)
	}
}