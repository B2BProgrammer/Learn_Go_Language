package main

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

/*
Lesson 11 :  Accessing the Internet - Web App Creation : News Aggregator

<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
*/

/*
[5]type == array // This is fixed size
[]type == slice // not fixed size
 */

type  SitemapIndex struct {
	Location []Location `xml:"sitemap"`
} 

type Location struct {
	Loc string `xml:"loc"`
} 


func main()  {

	// Parsing the XML from Internet
	resp,_ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") // Use Http & Get method to get the body
	bytes,_ := ioutil.ReadAll(resp.Body) // Read the resp body
	resp.Body.Close()


	var smp SitemapIndex
	xml.Unmarshal(bytes,&smp)

	fmt.Println(smp.Location)
	
}



