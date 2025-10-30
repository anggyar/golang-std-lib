package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Anggyar")
	data.PushBack("Muhamad")
	data.PushBack("Yahya")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Anggyar

	next := head.Next() // Muhamad
	fmt.Println(next.Value)

	next2 := next.Next() // Yahya
	fmt.Println(next2.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
