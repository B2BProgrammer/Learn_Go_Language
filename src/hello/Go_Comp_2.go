package main

import (
	"time"
	"fmt"
	"sync"
)

/*
 Concurrency Vs Parallel

Concurrency : Not a parallel, dealing with multiple things at a time
Parallel : DOING multiple things at a time

In Go : Concurrency is implemented using GO routine, Goroutine Synchronization with Concurrency :

NOTE NOTE :
if .. waitGroup.Done() .. is not executed .. then it will enter DEADLOCK situation
 ---- Try this by commenting // waitGroup.Done() : It will result in DEADLOCK situation

Important concepts to Realize Concurrency in Go

1. Go  : append "Go" to statment for implementing Synchronization
2. WaitGroup [Add, wait, Done]  : utilize Synchronization's waitgroup variable for communcication between method calls or Threads
3. Defer : append "Defer" - to make sure at any cost, this is LIFO [ last statment to be executed]
4.a Panic : Hold availble to programmer to stop the execution of program abruptly
4.b Recover  : In case of Panic, it is smoothly handled in recover(r) -- r := recover() .. if this is NOT null ..
5. Channels[Chan] : Channels, used along with GoRoutine to send & recieve values between channels using channel operator " <-"
 */




var wg sync.WaitGroup // WaitGroup, used to notify .. when the go routine is finished
func main(){

	fmt.Println("Start ------------------SECTION 1---------------------")
	wg.Add(1)
	go say("Hey")  // Added Go to make .. Test by removing "go"

	wg.Add(1)
	go say("There")

	wg.Wait()

	//time.Sleep(time.Second)
	fmt.Println("End ------------------SECTION 1---------------------")


	/*
	CHANNEL Understanding below
	 */
	fmt.Println("Start ------------------SECTION 2---------------------")
	// Make a channel : Chan type int, buffer for 10 items only
	fooVal1 := make(chan int,10 )

	for i:=0;i<10 ;i++  {
		wg.Add(1)
		go channelMethod(fooVal1,i)
	}
	wg.Wait()

	close(fooVal1) // Closing the channel

	// Enhanced for loop to iterate over the channel fooVal1
	for item := range fooVal1{
		fmt.Println(item)
	}
	time.Sleep(time.Second)
	fmt.Println("End ------------------SECTION 2---------------------")
}

func say(s string)  {
	//1.defer[Will make sure this is last statment to execute
	//2.Notify hte waitGRoup, that this routine is done, so send another thread to execute
	defer wg.Done()
	defer cleanup1()
	for i:= 0;i< 3 ;i++  {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
		if i == 2 {
			panic(" Oh panic it is 2")
		}
	}
}

func cleanup1(){
	if r:= recover(); r!= nil{
		fmt.Println("Recovered the cleanup")
	}
}


func channelMethod(channel1 chan int, value1 int){
	defer  wg.Done()
	channel1 <- value1*5
}