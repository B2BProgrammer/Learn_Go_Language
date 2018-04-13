package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

/*
Lesson 10 :  Accessing the Internet - Web App Creation : News Aggregator
*/

func main()  {

	// Parsing the XML from Internet
	resp,_ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") // Use Http & Get method to get the body
	bytes,_ := ioutil.ReadAll(resp.Body) // Read the resp body
	string_body := string(bytes) // Typcast bytees to String DataType
	fmt.Println(string_body) // Print the String - Which is the XML
	resp.Body.Close()  // Always close the response
}



