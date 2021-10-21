package main

import "fmt"

func main(){

	nama1 := "Reyner"
	nama2 := "Anton"

	//Is R > A ?
	result := nama1 > nama2
	fmt.Println(result)

	result = nama1 == nama2
	fmt.Println(result)

	value1 := 100
	value2 := 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}