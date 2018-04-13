package main

/*
Web development basis
Testing comment
 */

import (
	"fmt"
	"net/http"
)


func index_handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Hey there</h1>")

	fmt.Fprintf(w, "<h1>Hey there</h1>")
}

func main()  {

	http.HandleFunc("/",index_handler2)
	http.ListenAndServe(":8000", nil)
	
}