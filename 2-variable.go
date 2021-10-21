package main

import "fmt"

func main(){

	var hal = "Belajar Go Lang"
	fmt.Println(hal)
	fmt.Println(len(hal))
	fmt.Println(hal[0])
	
	//number
	//uint = 8,16,32,64
	//int = 8,16,32,64
	//float = 32,64
	//complex = 64,128 -> like float but have real n imaginary
	var a1 uint8 = 1
	var b1 uint8 = 2
	fmt.Println(a1+b1)

	var a2 int32 = -32000
	var b2 int32 = 64000
	fmt.Println(a2+b2)

	//boolean
	//var status bool = true

	//an easy way
	nama := "Reyner"
	umur := 22
	fmt.Println(nama,umur)

	var (
		firstName = "Reyner"
		lastName = "Pantou"
	)
	fmt.Println(firstName,lastName)
	
}