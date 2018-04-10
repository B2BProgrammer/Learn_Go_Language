package main

/*
Lesson 4 :  Pointers in Go Lang
x := 15  // variable Declartion
a := &x // memory address of x, stored at a
*a // accessing the value of x, using pointers
 */

import "fmt"

/*
First Method - main() to be called as part of Go Language Flow
 */
func main()  {
	x := 15  // variable Declartion
	a := &x // memory address of x

	fmt.Println("Value of x :",x) // Display value of x
	fmt.Println("Address of x :",&x) // Display memory address of x : 0xc042054080
	fmt.Println("Address of x, stored in variable :",a) // Display the stored memory address of x : 0xc042054080
	fmt.Println("Value of x, Referenced using memory address :",*a)

	*a = 5;
	fmt.Println("Value of x, Referenced using memory address after updating using pointers :", *a)

	*a = *a**a
	fmt.Println("Manipulation of value of variables using pointers :",x)
}



