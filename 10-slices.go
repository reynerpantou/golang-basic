package main

import "fmt"

func main(){

	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Changed"
	fmt.Println(slice1)

	slice1[0] = "MeiChanged"
	fmt.Println(slice1)
	fmt.Println(months)

	//taking value from index 10 until the end
	var slice2 = months[10:] 
	fmt.Println(slice2)

	//append some value into slice
	//when the len reach the maximum then recreate array & have no impact to the real array if anything changes
	var slice3 = append(slice2, "NewMonth")
	fmt.Println(slice3)

	slice3[0] = "NewNovember"
	fmt.Println(slice3)

	fmt.Println("The Actual Array =", months)
	
	//Create slice in safest way
	//make(type, length, capacity)
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Reyner"
	newSlice[1] = "Pantou"
	fmt.Println(newSlice)

	//Copy Slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	copySlice2 := make([]string, len(newSlice)-1, cap(newSlice))
	copy(copySlice2, newSlice)
	fmt.Println(copySlice2)

	//Difference format between array and slice
	thisArray := [5]int {1, 2, 3, 4, 5}
	thisArray2 := [...]int {1, 2, 3, 4, 5}

	//if theres no value or spread operator (...) then its supposed to be slice
	thisSlice := []int {1, 2, 3, 4, 5} 

	fmt.Println(thisArray)
	fmt.Println(thisArray2)
	fmt.Println(thisSlice)
}