package main

import (
	"fmt"
	"golang-basic/helper"
)

func main() {

	//Public function from `helper` package
	helper.SayHello("Reyner")

	//Public variable from `helper` package
	fmt.Println(helper.Application)
}
