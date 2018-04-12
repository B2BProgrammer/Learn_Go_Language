package main

import (
	"fmt"
	"time"
	"sync"
)

/*
Lesson 23 : Channels, used along with GoRoutine to send & recieve values between channels using channel operator " <-"

Buffering & Iterating over Channels
 */


 var wg sync.WaitGroup

func main()  {
	// Make a Channel : chan Type int
	fooVal := make(chan int, 10) // Buffer for 10 items

	for i:= 0; i < 10 ;i++  {
		wg.Add(1)
		go foo(fooVal,i)
	}

	wg.Wait()
	close(fooVal) // Closing the channel

	// Enhanced for looop
	for item := range fooVal{
		fmt.Println(item)
	}

	time.Sleep(time.Second)

}

// Pass : Channel & someValue
func foo(c chan int, someValue int)  {
	defer wg.Done()
	c <- someValue * 5

	// *** No need of return because data is sent/received using Channel
}