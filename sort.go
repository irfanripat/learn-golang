package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int
}

type Users []User

func (users Users) Len() int {
	return len(users)
}

func (users Users) Less(i, j int) bool {
	return users[i].Age < users[j].Age
}

func (users Users) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}

func main() {
	users := Users{
		User{"John", 20},
		User{"Jane", 30},
		User{"Joe", 25},
	}
	
	fmt.Println(users)
	
	sort.Sort(users)
	
	fmt.Println(users)
}