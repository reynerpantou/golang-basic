package main

import "fmt"

func main(){

	nilaiAkhir := 90
	absensi := 75

	lulusKelas := nilaiAkhir >= 80
	lulusAbsensi := absensi >= 80
	fmt.Println(lulusKelas)
	fmt.Println(lulusAbsensi)
	
	var lulus bool
	lulus = lulusKelas && lulusAbsensi
	fmt.Println(lulus)
}