package main

import "fmt"

func sumAll(numbers ...int) int {
	var total int
	for _, value := range numbers {
		total += value
	}
	return total
}
func main() {
	grade := sumAll()
	fmt.Println(grade)

	//Array
	gradeArr := sumAll(5, 7, 3, 2, 1, 3)
	fmt.Println(gradeArr)

	//Slice
	num := []int{100, 90, 60}
	//should add spread operator at the end of variable
	//to show the golang that we use slice
	gradeSlice := sumAll(num...)
	fmt.Println(gradeSlice)
}
