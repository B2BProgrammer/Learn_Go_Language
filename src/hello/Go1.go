package main

import ("fmt"
		"net/http" )

func  main()  {

/////////////////////Part 3 : Types//////////////////////////////////////////////

	// Number Addition
	num1,num2 := 5.6,3.1
	fmt.Println(add(num1,num2))

	// String Concatenation
	w1, w2 := "Hey", "Ajith"
	fmt.Println(stringConcatenate(w1,w2))

	// Number Type Conversion
	var  a int  = 62
	var b float64 = float64(a)
	fmt.Println(b)

////////////////////////// Part 4 : Pointers/////////////////////////////////////////

	x := 15 // What type is 15?
	addX := &x // memory address
	fmt.Println("Address of X in memory :")
	fmt.Println(addX) // hex value
	fmt.Println(*addX)  // 15
	*addX = 5
	fmt.Println(x)  // x value is changed to 5, accesing the memory & overwriting the value

	*addX = *addX * *addX
	fmt.Println(*addX)


	////////////////////////// Part 5 : Simple Web Server/////////////////////////////////////////

	// handlers
	//http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)


	////////////////////////// Part 6 : Struts (Equivalent of Class)/////////////////////////////////////////

	



	} // End - main

func  add(x ,y float64) float64 {
	return x+y
}

func stringConcatenate(a,b string) (string,string) {
	return a,b
}

func index_handler(w http.ResponseWriter, r *http.Request ){
	fmt.Fprintf(w, "whoa, Go is neat!")
}

func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Amazing about handler")
}


