package main

import (
	"fmt"

)

/*
Lesson 20 : Defer

Defer in "GO" ..  if .. waitGroup.Done() .. is not executed .. then it will enter DEADLOCK situation
 ---- Try this by commenting // waitGroup.Done() : It will result in DEADLOCK situation



 */

func main()  {
 foo()
 foo1()
}


/*
defer statements are executed at last in the method, after all teh statments are executed ..

Eg : sequence of execution
1. Doing some stuff, Who knows what
2. Are we Done
3. Done !!
 */
func foo(){
	defer fmt.Println("Done !!")
	defer fmt.Println("Are we done")
	fmt.Println("Doing some stuff, Who knows what")
}


/*
defer statements are executed at last in the method, after all teh statments are executed ..

LIFO - Last in First Out
Eg : sequence of execution
4
3
2
1
0
 */

func foo1(){
	for i:=0 ; i <5 ; i++ {
		defer fmt.Println(i)
	}
}