package main

import (
	"fmt"
	"strconv"
)

func main(){

	//String to Boolean
	boolean, err := strconv.ParseBool("false")
	if err == nil {
		fmt.Println("Boolean :", boolean)
	}else{
		fmt.Println("Error :", err.Error())
	}

	//String to Int : string_number, base_int(2, 8, 10, 16), bitsize(32, 64, etc) 
	number, err := strconv.ParseInt("1000000", 10, 32)
	if err == nil {
		fmt.Println("Number :", number)
	}else{
		fmt.Println("Error :", err.Error())
	}
	
	//Int to String : int, base_int(2, 8, 10, 16)
	number2 := strconv.FormatInt(2560, 16)
	fmt.Println("Number2: ", number2)

	number3, err := strconv.Atoi("456")
	if err == nil {
		fmt.Println("Number3 :", number3)
	}else{
		fmt.Println("Error :", err.Error())
	}

	number4 := strconv.Itoa(789)
	fmt.Println("Atoi :", number4)

	number5, err := strconv.Atoi("789")
	if err == nil{
		fmt.Println("Itoa :", number5)
	}else{
		fmt.Println("Error :", err.Error())
	}
}