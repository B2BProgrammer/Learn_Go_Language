package main

/*
Lesson 3 :  Is typing Information, Go has built-in WebServer
Data Types :
bool
int
String
 */

import (
	"fmt"  //
)

/*
First Method - main() to be called as part of Go Language Flow
 */
func main()  {
	// Declaration Eg : 1
	var num1 float64 = 10.0;
	var num2 float64 = 9.5;
	fmt.Println(add(num1,num2))

	// Declaration Eg : 2 - mulitple variables can be declared together
	var num3, num4 float64 = 10,2;
	fmt.Println(add1(num3,num4))

	// Declaration Eg : 3 - := syntax can be used to mulitple variables, without DataType Declaration
	num5, num6 := 11.0, 11.0;
	fmt.Println(add1(num5,num6))

	// Declaration Eg : 4 [Constant]
	const x int  = 5;

	// Declaration Eg : 5 [Strings]
	w1, w2 := "Hey", "There"
	fmt.Println(multiple(w1,w2))

	// Declaration Eg : 6 [
	var a int = 62
	var b float64 = float64(a)

	x:= a // x will be type int
}

/*
method/function signature :
 */
func add(x float64, y float64) float64{
	return x + y
}

/*
return single output
 */
func add1(x,y float64) float64  {
	return x+y
}

/*
 Return : multiple output
Signature Eg :

func funcName( parameters) (return1, return2, return 3 .. )
 */
func  multiple(a,b string) (string,string)  {
	return a,b
}



