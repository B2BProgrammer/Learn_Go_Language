package main

import "fmt"

/*
Lesson 14 : DataStructure [MAPS]

 */
func main()  {
  // i.e k-v : [String]float32
 //var grades map[String]float32

 // Declartion of grades Map, having K-v map[string]float32
 //Make - will make up a maps
 grades := make(map[string]float32)

 // 1. Adding K-V to Map
 grades["Ajith"] = 42
 grades["Divya"] = 54
 grades["Megha"] = 95
 grades["Shobha"] = 78
 grades["Chandru"] = 55

 // 2. Printing the grades Complete Map
 fmt.Println(grades)

 // 3. Getting the Value from the grades Map
 AjithGrade := grades["Ajith"]
 fmt.Println(AjithGrade)

 // 4. Deleteing the K-V from the Map
 delete(grades,"Divya")
 fmt.Println(grades)

 //5. Iteration over the Map
 for key,value := range grades{
 	fmt.Println("KeyName :", key,"-","Value: ", value)
 }

 //6


}