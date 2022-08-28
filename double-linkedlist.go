package main

import (
	"fmt"
	"container/list"
)

func main() {
	data := list.New()
	data.PushBack("Irfan")
	data.PushBack("Ripat")

	for e:= data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}