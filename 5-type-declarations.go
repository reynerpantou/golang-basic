package main

import "fmt"

func main() {

	//make an alias using existing data type
	type KTP string
	type Married bool

	var reyner KTP = "12395850394854"
	var status Married = false
	fmt.Println(reyner, status)
}
