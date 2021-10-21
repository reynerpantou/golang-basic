package main

import (
	"fmt"
	"flag"
)

func main(){

	//go run 38-flag.go -hostname=localhost -username=reyner -password=reyner99

	//Pattern (parameter, default_value, keyword `up to you`)
	hostname := flag.String("hostname", "localhost", "Enter your hostname")
	username := flag.String("username", "root", "Enter your username")
	password := flag.String("password", "root", "Enter your password")

	//Need to parse or the flag will return the default value
	flag.Parse()

	fmt.Println("Hostname :", *hostname)
	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)
}