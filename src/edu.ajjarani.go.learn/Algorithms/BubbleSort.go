package Algorithms

import "fmt"

func main() {

	/*
	// Arrange in an Ascending order

	5,62,55,7,9,33

	5,62,55,7,9,33 [1st Itter]
i	5,55,62,7,9,33 [2st Itter]
	5,55,7,62,9,33 [3st Itter]
	5,55,7,9,62,33 [3st Itter]
	5,55,7,9,33,62 [3st Itter]

	--------------
	5,55,7,9,33,62 [3st Itter]
	5,7,55,9,33,62
	5,7,9,55,33,62
	5,7,9,33,55,62
	5,7,9,33,55,62

	---------------------
	5,7,9,33,55,62

	*/
	// 5,7,9,33,5,62

	var array1[7]int
	array1[0] = 5
	array1[1] = 62
	array1[2] = 55
	array1[3] = 7
	array1[4] = 9
	array1[5] = 33

	var arrary2[7]int = bubbleSort(array1)
	fmt.Println(arrary2)
}

func bubbleSort(array1 [7]int)  ([7]int) {
	for i:= 0;i < 5 ;i++  {
		for j:=0;j<5 ;j++  {
			if array1[j] > array1[j+1]{
				// swap
				/*
				temp := array1[j]
				array1[j] = array1[j+1]
				array1[j+1] = temp
				*/

				array1[j],array1[j+1] = array1[j+1],array1[j]
			}
		}
	}
	fmt.Println(array1)
 return  array1
}



