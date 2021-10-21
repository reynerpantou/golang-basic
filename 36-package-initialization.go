package main

import (
	"fmt"
	"golang-basic/database"

	//use underscore A.K.A Blank Identifier to bypass the package that we dont specificly used
	_ "golang-basic/helper"
)

func main() {
	db := database.GetDatabase()
	fmt.Println(db)
}
