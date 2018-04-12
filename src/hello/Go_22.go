package main

import "fmt"

/*
Lesson 22 : Channels, used along with GoRoutine to send & recieve values between channels using channel operator " <-"

 */

func main()  {
	// Make a Channel : chan Type int
	fooVal := make(chan int)

	go foo(fooVal, 5)
	go foo(fooVal, 3)

	v1 := <- fooVal
	v2 := <- fooVal

	fmt.Println("v1 is recived from foo, using channels NOT return statment :",v1)
	fmt.Println("v2 is recived from foo, using channels NOT return statment :",v2)

}

// Pass : Channel & someValue
func foo(c chan int, someValue int)  {
	c <- someValue * 5

	// *** No need of return because data is sent/received using Channel
}