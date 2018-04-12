package main

import (
	"fmt"
	"time"
)

/*
Lesson 18 : Concurrency :

note : Concurrency Vs Parallel

Concurrency : Not a parallel, dealing with multiple things at a time
Parallel : DOING multiple things at a time


In Go : Concurrency is implemented using GO routine
 */
func main()  {
 go say("Hey")  // Added Go to make .. Test by removing "go"
 go say("There")

 time.Sleep(time.Second)

 /*

  */
}

func say(s string)  {
	for i := 0 ; i <3 ; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}
}