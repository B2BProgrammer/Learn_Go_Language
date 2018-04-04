package main

/*
METHODS
2 kinds of methods
1. Value receivers : Just get the value
2. Pointer receivers : Change the value
 */


import "fmt"

 const usixteenbitmax float64 = 65535
 const kmh_multiple float64  = 1.60934

// Equivalent to class/Declartion of car class
type car1 struct {
	gas_pedal uint16 //min vlau = 0, max - 65535
	brake_pedal uint16
	steering_wheel int16 // -32k to +32k
	top_speed_kmh float64
}

/*
Method Eg : Value receivers
 */
func (c car1)  kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}

/*
Method Eg : Value receivers
 */
func (c car1)  mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/kmh_multiple)
}


func main()  {
	// Equivalent to object instatiate : creation of a_car_1 object
	// Equivalent to object instatiate : creation of a_car_2 object
	a_car :=car1{22343,2,233, 100}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}


