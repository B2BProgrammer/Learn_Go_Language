package main

import "fmt"

/*
Lesson 7 :  METHODS in Go
This chapter is about : Value Receivers

2 kinds of methods
1. Value receivers : Just get the value
2. Pointer receivers : Change the value
*/

const usixteenbitmax float64 = 65535
const kmh_multiple float64  = 1.60934

// Equivalent to class/Declartion of car class
type car7 struct {
	gas_pedal uint16 //min value = 0, max - 65535
	brake_pedal uint16
	steering_wheel int16 // -32k to +32k
	top_speed_kmh float64
}

/*
Method Eg : Value receivers
Input paramater : c - object, DataType : car7
c, object will have all the values, required .. not necessary to explicity pass each one
 */
func (c car7)  kmh() float64 {
	fmt.Println("Value Recieved from Caller :",c.brake_pedal,c.gas_pedal,c.steering_wheel,c.top_speed_kmh)
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}


/*
Method Eg : Value receivers
 */
func (c car7)  mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/kmh_multiple)
}


func main()  {
	// Equivalent to object instatiate : creation of a_car_1 object
	a_car :=car7{223,2,233, 100}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh()) // call the method kmh(), NOTE : varibale values are not explictily passed, just they are accessed
	fmt.Println(a_car.mph()) // call the method mph()
}
