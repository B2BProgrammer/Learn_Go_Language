package main

/*
Lesson 5 :  Creation of Webserver

After success, running : access below URL
http://localhost:9000/
http://localhost:9000/about/

/*
First Method - main() to be called as part of Go Language Flow
 */

import (
	"net/http"
	"fmt"
)
func main()  {
	// Using net/http library, call the index_handler
	http.HandleFunc("/", index_handler)  //index_handler, this is library function
	http.HandleFunc("/about/", about_handler) //about_handler, this is library function
	http.ListenAndServe(":9000",nil)
}


func index_handler(w http.ResponseWriter, r *http.Request)  {
 fmt.Fprintf(w, "wow, Go is wonderfull!")
}

func about_handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Show me the about page !!")
}



