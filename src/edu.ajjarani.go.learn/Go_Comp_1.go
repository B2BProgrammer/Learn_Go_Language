package main

import (
	"fmt"
	"math/rand"
)


// Equivalent to class/Declartion of lexus class
type lexus struct {
	gas_pedal uint16 //min vlau = 0, max - 65535
	brake_pedal uint16
	steering_wheel int16 // -32k to +32k
	top_speed_kmh float64
}


/*
Final Go Lang Code, Which utilizes all hte concepts available in Go Lang
 */
func main(){

	// 1. Using fmt
	fmt.Println("START--------------SECTION 1-----------------------")
	fmt.Println("Start of Go Execution Flow is always from main()")
	fmt.Println("END--------------SECTION 1-----------------------")


	//2. Call of methods
	fmt.Println("START--------------SECTION 2-----------------------")
	foo2()
	fmt.Println("END--------------SECTION 2-----------------------")

	/*
	//3. DataType & Variable Declaration
	  Primitive : 1.bool, 2.signed & 3.unsigned int , 4. String
	 */
	// 3.a Declaration Eg : 1
	fmt.Println("START--------------SECTION 3-----------------------")
	var num1 float64 = 10.0;
	var num2 float64 = 9.5;
	fmt.Println(add(num1,num2))

	// 3.b Declaration Eg : 2 - mulitple variables can be declared together
	var num3, num4 float64 = 10,2;
	fmt.Println(add1(num3,num4))

	//  3.c Declaration Eg : 3 - := syntax can be used to mulitple variables, without DataType Declaration
	num5, num6 := 11.0, 11.0;
	fmt.Println(add1(num5,num6))

	//  3.d Declaration Eg : 4 [Constant]
	const x int  = 5;

	//  3.e Declaration Eg : 5 [Strings]
	w1, w2 := "Hey", "There"
	fmt.Println(multiple(w1,w2))

	//  3.f Declaration Eg : 6 [Typecasting]
	var a int = 62
	var aTypeCasted float64 = float64(a) // a
	fmt.Println("TypeCasted Value of a", aTypeCasted)


	/*
	  3.g for Loop & if Statments Example
	  */
	for j :=0; j< 10 ;j++  {
		if j < 5 {
			fmt.Println("The Value of j below 5 :",j)
		} else {
			fmt.Println("The Value of j above 5 :",j)
		}
	}
	fmt.Println("END--------------SECTION 3-----------------------")

	/*
	  4. Pointers understanding in Go Lang
	 */
	fmt.Println("START--------------SECTION 4-----------------------")
	varPoint := 15  // variable Declartion
	varPointMemAdd := &varPoint // memory address of x

	fmt.Println("Value of varPoint :",varPoint) // Display value of x
	fmt.Println("Address of varPoint :",&varPoint) // Display memory address of x : 0xc042054080
	fmt.Println("Address of varPoint, stored in variable :",varPointMemAdd) // Display the stored memory address of x : 0xc042054080
	fmt.Println("Value of varPoint, Referenced using memory address :",*varPointMemAdd)

	*varPointMemAdd = 5
	fmt.Println("Value of varPoint, Referenced using memory address after updating using pointers :", *varPointMemAdd)
	*varPointMemAdd = *varPointMemAdd**varPointMemAdd
	fmt.Println("Manipulation of value of variables using pointers :",varPoint)
	fmt.Println("START--------------SECTION 4-----------------------")

	/*
	 5.	// Equivalent to object instatiate : creation of lexus_1 & lexus_2 object
	 */
	fmt.Println("START--------------SECTION 5-----------------------")
	lexus_1 := lexus{ gas_pedal: 22341, brake_pedal: 0, steering_wheel: 12561, top_speed_kmh: 225.0}
	lexus_2 :=lexus{22343,2,233, 100}

	fmt.Println(lexus_1.gas_pedal)
	fmt.Println(lexus_2.top_speed_kmh)
	fmt.Println(lexus_2.kmh())  // Accesing Value reciver method
	lexus_2.new_top_speed(500)    // Accesing pointer reciver
	fmt.Println("START--------------SECTION 5-----------------------")


	/*
	 6. DataStrutures in Go Lang
		a. MAPS
	 */
	// i.e k-v : [String]float32
	//var grades map[String]float32

	// Declartion of grades Map, having K-v map[string]float32
	//Make - will make up a maps
	fmt.Println("START--------------SECTION 6-----------------------")
	Scores := make(map[string]float32)

	// 1. Adding K-V to Map
	Scores["Ajith"] = 42
	Scores["Divya"] = 54
	Scores["Megha"] = 95
	Scores["Shobha"] = 78
	Scores["Chandru"] = 55

	// 2. Printing the Scores Complete Map
	fmt.Println(Scores)

	// 3. Getting the Value from the grades Map
	AjithGrade := Scores["Ajith"]
	fmt.Println(AjithGrade)

	// 4. Deleteing the K-V from the Map
	delete(Scores,"Divya")
	fmt.Println(Scores)

	//5. Iteration over the Collection[Map] : Enhanced For Loop
	for key,value := range Scores{
		fmt.Println("KeyName :", key,"-","Value: ", value)
	}
	fmt.Println("START--------------SECTION 6-----------------------")

}





/*
Func foo2, called from main
 */
func foo2(){
	fmt.Println("foo2, called from main()")
	fmt.Println("A random number generation :", rand.Intn(100))  // Using standard library of Go Lang
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


/*
Method Eg : Value receivers
Input paramater : c - object, DataType : lexus
c, object will have all the values, required .. not necessary to explicity pass each one
 */
func (l lexus)  kmh() float64 {
	fmt.Println("Value Recieved from Caller :",l.brake_pedal,l.gas_pedal,l.steering_wheel,l.top_speed_kmh)
	return float64(l.gas_pedal) * (l.top_speed_kmh/10)

	// Note note : Value reciver .. need a return statment, as value is passed from method 1 to method2
}


/*
Eg : Pointer method :
Changing the value in the class

If Strut is big, the Pointer method is good/cheap
Less garbage usage, So better to pointer is better
*/
func (l *lexus) new_top_speed (newspeed float64) {
	l.top_speed_kmh = newspeed

	// Note note : pointer reciver .. DON'T need a return statment, as value is  NOT passed from method 1 to method2
	// Only memory pointers are used to manipulate the value
}