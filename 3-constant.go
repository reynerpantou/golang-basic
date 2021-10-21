package main

import "fmt"

func main(){
	//const not showing error even not used in the code
	const status bool = true
	const check = false

	fmt.Println(status, check)

	const(
		firstName = "Reyner"
		lastName = "Pantou"
	)

	fmt.Println(firstName, lastName)
}