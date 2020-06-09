package main

import "fmt"

type node struct {
	prev  *node
	next  *node
	value interface{}
}

type linkedList struct {
	head   *node
	tail   *node
	length int
}

func new(value interface{}) linkedList {
	ll := linkedList{
		head: newNode(value),
	}
	ll.tail = ll.head
	ll.length = 1
	return ll
}

func newNode(value interface{}) *node {
	return &node{
		next:  nil,
		prev:  nil,
		value: value,
	}
}

func (ll *linkedList) append(value interface{}) {
	newNode := newNode(value)
	newNode.prev = ll.head
	ll.tail = newNode
	ll.head.next = newNode
	ll.length++
}

func (ll *linkedList) prepend(value interface{}) {
	newNode := newNode(value)
	newNode.next = ll.head
	ll.head = newNode
	ll.length++

}

func (ll *linkedList) insert(index int, value interface{}) {
	if index == 0 {
		ll.prepend(value)
		return
	}
	newNode := newNode(value)
	previousNode := ll.traverse(index - 1)
	newNode.prev = previousNode
	newNode.next = previousNode.next
	previousNode.next = newNode
}

func (ll *linkedList) remove(index int) {
	previousNode := ll.traverse(index - 1)
	previousNode.next = previousNode.next.next
}

func (ll *linkedList) traverse(index int) *node {

	head := ll.head
	if index >= ll.length {
		index = ll.length - 1
	}
	for i := 0; i < index; i++ {
		head = head.next
	}

	return head
}

func (ll *linkedList) print() {

	head := ll.head
	if ll.length == 1 {
		fmt.Println(head.value)
		return
	}
	for head != nil {
		fmt.Println(head.value)
		head = head.next
	}
}

func main() {
	myll := new(10)
	myll.append(15)
	myll.prepend(5)
	myll.insert(1, 22)
	myll.insert(0, 99)
	myll.print()
	myll.remove(1)
	myll.print()
}
