package main

import "fmt"

func endApp() {
	fmt.Println("Application closed")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Application error")
	}
	fmt.Println("Application running")
}
func main() {
	runApp(true)

}
