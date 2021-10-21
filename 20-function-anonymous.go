package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) == true {
		fmt.Println("You are Blocked,", name)
	} else {
		fmt.Println("Hello,", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	//best way
	registerUser("admin", blacklist)
	registerUser("reyner", blacklist)

	//redundant way
	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("admin", func(name string) bool {
		return name == "root"
	})
}
