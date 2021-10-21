package main

import "fmt"

func main(){

    var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)

	//overflow
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	//String Conversion

	var kata = "ABC"

	//Automatically set to byte because the ascii number < 128, in this case only
	//var ascii = huruf[0] OR ascii := huruf[0]
	var ascii byte = kata[0]
	
	var huruf = string(ascii)
	fmt.Println(huruf)

}
	