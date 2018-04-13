package main

import "fmt"

/*
Lesson 6 :  Creation of STRUTS [

Equivalent of Class in Go Lang
*/

// Equivalent to class/Declartion of car class
type car struct {
	gas_pedal uint16 //min vlau = 0, max - 65535
	brake_pedal uint16
	steering_wheel int16 // -32k to +32k
	top_speed_kmh float64
}



func main()  {
	// Equivalent to object instatiate : creation of a_car_1 object
	a_car_1 := car{ gas_pedal: 22341, brake_pedal: 0, steering_wheel: 12561, top_speed_kmh: 225.0}

	// Equivalent to object instatiate : creation of a_car_2 object
	a_car_2 :=car{22343,2,233, 100}


	fmt.Println(a_car_1.gas_pedal)
	fmt.Println(a_car_2.top_speed_kmh)
}



