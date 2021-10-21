package main

import "fmt"

//We can use empty interface if there are several cases that needs any parameter or any return
//for any parameter we can use "func Oops(data interface{}) interface{}"
func Oops(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Oops"
	}
}

func main() {

	//The original
	//var empty interface{} = Oops(1)

	//Short version
	empty := Oops(2)

	//Error
	//var data int = Oops(1)
	//even Oops(1) return is integer, its not gonna work
	//we cant set specific data type for empty interface

	fmt.Println(empty)
}
