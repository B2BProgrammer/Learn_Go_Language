package main

/*
METHODS
2 kinds of methods
1. Value receivers : Just get the value
2. Pointer receivers : Change the value

Pointer receivers : Eg Here
 */


import "fmt"

 const usixteenbitmax float64 = 65535
 const kmh_multiple float64  = 1.60934

// Equivalent to class/Declartion of car class
type car8 struct {
	gas_pedal uint16 //min vlau = 0, max - 65535
	brake_pedal uint16
	steering_wheel int16 // -32k to +32k
	top_speed_kmh float64
}

/*
Method Eg : Value receivers
 */
func (c car8)  kmh() float64 {
	c.top_speed_kmh = 500
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}

/*
Method Eg : Value receivers
 */
func (c car8)  mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/kmh_multiple)
}


func newer_top_speed(c car, speed float64) car {
	c.top_speed_kmh = speed
	return  c
}

/*
Eg : Pointer method :
Changing the value in the class

Less garbage usage, So better to pointer is better
 */
func (c *car8) new_top_speed (newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main()  {
	// Equivalent to object instatiate : creation of a_car_1 object
	// Equivalent to object instatiate : creation of a_car_2 object
	a_car :=car8{22343,2,233, 100}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	a_car.new_top_speed(500)  // Pointer method access : changing the values
	fmt.Println(a_car.kmh())  // Value Method Access : Getting the values
	fmt.Println(a_car.mph())
	fmt.Println(a_car.top_speed_kmh)

}


