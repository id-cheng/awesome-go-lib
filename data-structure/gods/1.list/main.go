package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/lists/doublylinkedlist"
	"github.com/emirpasic/gods/lists/singlylinkedlist"
	"github.com/emirpasic/gods/utils"
)

func main() {
	newArrayList()
	newSinglyLinkedList()
	newDoublyLinkedList()
}

func newArrayList() {
	list := arraylist.New()
	list.Add(1, 10, 5)
	list.Sort(utils.IntComparator)
	fmt.Println(list.String())
	val, hasVal := list.Get(0)
	fmt.Println(val, hasVal)
	val, hasVal = list.Get(100)
	fmt.Println(val, hasVal)
	list.Swap(1, 2)
	fmt.Println(list.String())
	list.Remove(1)
	fmt.Println(list.Empty(), list.Size())
	list.Clear()
	fmt.Println(list.Empty(), list.Size())
	list.Insert(0, 1, 2, 3)
	fmt.Println(list.String())
}

func newSinglyLinkedList() {
	list := singlylinkedlist.New()
	list.Add(1, 10, 5)
	list.Sort(utils.IntComparator)
	fmt.Println(list.String())
	val, hasVal := list.Get(0)
	fmt.Println(val, hasVal)
	val, hasVal = list.Get(100)
	fmt.Println(val, hasVal)
	list.Swap(1, 2)
	fmt.Println(list.String())
	list.Remove(1)
	fmt.Println(list.Empty(), list.Size())
	list.Clear()
	fmt.Println(list.Empty(), list.Size())
	list.Insert(0, 1, 2, 3)
	fmt.Println(list.String())
}

func newDoublyLinkedList() {
	list := doublylinkedlist.New()
	list.Add(1, 10, 5)
	list.Sort(utils.IntComparator)
	fmt.Println(list.String())
	val, hasVal := list.Get(0)
	fmt.Println(val, hasVal)
	val, hasVal = list.Get(100)
	fmt.Println(val, hasVal)
	list.Swap(1, 2)
	fmt.Println(list.String())
	list.Remove(1)
	fmt.Println(list.Empty(), list.Size())
	list.Clear()
	fmt.Println(list.Empty(), list.Size())
	list.Insert(0, 1, 2, 3)
	fmt.Println(list.String())
}
