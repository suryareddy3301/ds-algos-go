package main

import "fmt"

type singlyNode struct {
	next  *singlyNode
	value interface{}
}

type singlyLinkedList struct {
	head   *singlyNode
	tail   *singlyNode
	length int
}

func newSNode(value interface{}) *singlyNode {
	return &singlyNode{
		next:  nil,
		value: value,
	}
}

func NewSinglyLinkedList(value interface{}) singlyLinkedList {
	node := newSNode(value)
	ll := singlyLinkedList{
		head:   node,
		length: 1,
	}
	ll.tail = ll.head
	return ll
}

func (ll *singlyLinkedList) append(value interface{}) {
	newNode := newSNode(value)
	ll.tail.next = newNode
	ll.tail = newNode
	ll.length++
}

func (ll *singlyLinkedList) prepend(value interface{}) {
	newNode := newSNode(value)
	newNode.next = ll.head
	ll.head = newNode
	ll.length++
}

func (ll *singlyLinkedList) insert(index int, value interface{}) {
	newNode := newSNode(value)
	previousNode := ll.traverse(index - 1)
	newNode.next = previousNode.next
	previousNode.next = newNode
	ll.length++
}

func (ll *singlyLinkedList) remove(index int) {
	previousNode := ll.traverse(index - 1)
	previousNode.next = previousNode.next.next
	ll.length--
}

func (ll *singlyLinkedList) traverse(index int) *singlyNode {
	node := ll.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (ll *singlyLinkedList) print() {
	head := ll.head
	for head != nil {
		fmt.Println(head.value)
		head = head.next
	}
}
