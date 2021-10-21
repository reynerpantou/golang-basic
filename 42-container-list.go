package main

import (
	"fmt"
	"container/list"
)

func main(){

	//Implement of Double Linked List
	data := list.New()
	
	data.PushBack("Reyner")
	data.PushBack("Pantou")

	//Front to Back
	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	//Back to Front
	for i := data.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}
}