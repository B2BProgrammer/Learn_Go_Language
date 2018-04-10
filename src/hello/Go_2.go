package main

/*
Lesson 2 :  A simple lesson to demonstrate the flow the Go lanng

1. Start of the Go Language code is from main(), & all methods should be called from main itself
2. math/rand [ syntax for importing the library]
 */

import (
	"fmt"  //
	"math" // Importing library math
	"math/rand"  // Importing library math/rand
)


func main()  {
	fmt.Println("The square root of 4 is ", math.Sqrt(4))
	foo()
}

/*
Helper method called from the main()
Syntax for calling the library methods :
// pacakage.Capital Letter
// Eg
// a. fmt.Println
// b. math.Sqrt
 */
func foo(){
	fmt.Println(" This foo() method - called from main()")
	fmt.Println(" A random number between 1-100 :", rand.Intn(100))
}