package linked_list

import "fmt"

type LinkedList struct {
	head *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (linkedList *LinkedList) InsertAtBeginning(value int) {
	newNode := &Node{
		value: value,
	}

	newNode.next = linkedList.head
	linkedList.head = newNode
	linkedList.size++
}

func (linkedList *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{
		value: value,
	}

	if linkedList.head == nil {
		linkedList.head = newNode
		linkedList.size++
		return
	}

	current := linkedList.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
	linkedList.size++
}

func (linkedList *LinkedList) Remove(value int) bool {

	if linkedList.head == nil {
		return false
	}

	if linkedList.head.value == value {
		linkedList.head = linkedList.head.next
		linkedList.size--
		return true
	}

	prev := linkedList.head
	current := linkedList.head.next
	for current != nil {
		if current.value == value {
			prev.next = current.next
			linkedList.size--
			return true
		}
		prev = current
		current = current.next
	}

	return false
}

func (linkedList *LinkedList) Search(value int) bool {
	current := linkedList.head
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}

	return false
}

func (linkedList *LinkedList) Print() {
	current := linkedList.head
	for current != nil {
		fmt.Printf("%d", current.value)
		if current.next != nil {
			fmt.Print(" -> ")
		}
		current = current.next
	}
	fmt.Println("")
}

func (linkedList *LinkedList) Size() int {
	return linkedList.size
}
