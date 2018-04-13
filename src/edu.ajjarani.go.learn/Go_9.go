package main

import (
	"net/http"
	"fmt"
)

/*
Lesson 9 :  WEB SERVER

Access the URL :
http://localhost:9001/

*/

func main()  {
  http.HandleFunc("/",index_handler)
  http.ListenAndServe(":9001",nil)
}

func index_handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"<h1>Hey There, Thanks for visiting </h1>")
	fmt.Fprintf(w,"<h1>GO Lang is super Fast</h1>")
	fmt.Fprintf(w,"<p>Thank you </p>")
}


