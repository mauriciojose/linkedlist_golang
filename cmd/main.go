package main

import (
	"fmt"
	linked_list "linkedlist/internal"
)

func main() {

	list := linked_list.NewLinkedList()

	list.InsertAtBeginning(3)
	list.InsertAtBeginning(1)
	list.InsertAtEnd(5)
	list.InsertAtEnd(7)
	list.InsertAtEnd(8)
	list.InsertAtEnd(9)
	list.InsertAtEnd(10)

	list.Print() // 1 -> 3 -> 5 -> 7 -> 8 -> 9 -> 10

	fmt.Println("Busca 5:", list.Search(5))
	fmt.Println("Remove 3:", list.Remove(3))
	fmt.Println("Remove 10:", list.Remove(10))

	list.Print()
	fmt.Println("Size:", list.Size())
}
