package main

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args

	//go run 37-os.go Reyner Pantou
	fmt.Println("Argument :")
	fmt.Println("[0]", args[0])
	fmt.Println("[1]", args[1])
	fmt.Println("[2]", args[2])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", name)
	}else{
		fmt.Println("Error :", err.Error())
	}

	//export GOLANG-USER=reyner
	//export GOLANG-PASS=reyner123
	user := os.Getenv("GOLANG_USER")
	pass := os.Getenv("GOLANG_PASS")

	fmt.Println("Username :", user)
	fmt.Println("Password :", pass)

}