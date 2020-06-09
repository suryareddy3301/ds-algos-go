package main

import "fmt"

type node struct {
	prev  *node
	next  *node
	value interface{}
}

type linkedList struct {
	head *node
	tail *node
}

func new(value interface{}) linkedList {
	ll := linkedList{
		head: &node{
			value: value,
			prev:  nil,
			next:  nil,
		},
	}
	ll.tail = ll.head
	return ll
}

func (ll *linkedList) print() {
	current := ll.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func (ll *linkedList) append(value interface{}) {
	node := newNode(value)
	ll.tail.next = node
	ll.tail = node
}

func (ll *linkedList) prepend(value interface{}) {
	node := newNode(value)
	node.next = ll.head
	ll.head = node
}

func newNode(value interface{}) *node {
	return &node{
		prev:  nil,
		next:  nil,
		value: value,
	}
}

func (ll *linkedList) insert(index int, value interface{}) {
	newNode := newNode(value)
	node := ll.head
	for i := 0; i < index-1; i++ {
		node = node.next
	}
	newNode.next = node.next
	node.next = newNode
}

func main() {
	list := new(10)
	list.append(30)
	list.prepend(5)
	list.insert(2, 15)
	list.insert(3, 20)
	list.insert(4, 25)
	// 5 10 15 20 25 30
	list.print()
}
