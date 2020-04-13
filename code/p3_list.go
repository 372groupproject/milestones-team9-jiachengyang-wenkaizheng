package main

import (
	"container/list"
	"fmt"
)

func listOperation() {
	List := list.New()
	List.PushBack("one")
	List.PushFront("zero")
	List.PushBack("two")
	element := List.PushBack("four")
	List.InsertAfter("five", element)
	List.InsertBefore("three", element)
	remove := List.PushBack("delete")
	List.Remove(remove)
	for i := List.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
func main() {
	listOperation()
}
