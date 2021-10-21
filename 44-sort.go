package main

import (
	"sort"
	"fmt"
)
type User struct{
	Name string
	Age int
}

type UserSlice []User

func (userSlice UserSlice) Len() int{
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main(){
	users := []User{
		{"Reyner", 22},
		{"Antonius", 40},
		{"Edward", 25},
		{"Franky", 23},
	}
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}