package main

import "fmt"

type doublyNode struct {
	prev  *doublyNode
	next  *doublyNode
	value interface{}
}

type doublyLinkedList struct {
	head   *doublyNode
	tail   *doublyNode
	length int
}

func newDNode(value interface{}) *doublyNode {
	return &doublyNode{
		prev:  nil,
		next:  nil,
		value: value,
	}
}

func newDoublylinkedList(value interface{}) doublyLinkedList {
	node := newDNode(value)

	ll := doublyLinkedList{
		head:   node,
		length: 1,
	}
	ll.tail = ll.head

	return ll
}

func (ll *doublyLinkedList) append(value interface{}) {
	newNode := newDNode(value)
	newNode.prev = ll.tail
	ll.tail.next = newNode
	ll.tail = newNode
}

func (ll *doublyLinkedList) prepend(value interface{}) {
	newNode := newDNode(value)
	ll.head.prev = newNode
	newNode.next = ll.head
	ll.head = newNode
}
func (ll *doublyLinkedList) traverse(index int) *doublyNode {
	node := ll.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}
func (ll *doublyLinkedList) insert(index int, value interface{}) {
	newNode := newDNode(value)
	previousNode := ll.traverse(index - 1)
	newNode.next = previousNode.next
	previousNode.next = newNode
}

func (ll *doublyLinkedList) remove(index int) {
	previousNode := ll.traverse(index - 1)
	previousNode.next = previousNode.next.next
}

func (ll *doublyLinkedList) reverse() {
	first := ll.head
	second := first.next
	for second != nil {
		temp := second.next
		//	second.prev = second.next
		second.next = first
		first = second
		second = temp
	}
	ll.head.next = nil
	ll.head = first
}

func (ll *doublyLinkedList) print() {
	node := ll.head

	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

//question :- reverse linked list in segments k
//eg.g input 1-2-3-4-5-6-7-8-9 k = 3
//output 3-2-1-6-5-4-9-8-7

func (ll *doublyLinkedList) reverseSegmented(k int, head *doublyNode) *doublyNode {
	current := head
	var prev, next *doublyNode

	count := 0

	/* Reverse first k nodes of linked list */
	for count < k && current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
		count++
	}

	/* next is now a pointer to (k+1)th node
	   Recursively call for the list starting from current.
	   And make rest of the list as next of first node */
	if next != nil {
		head.next = ll.reverseSegmented(k, next)
	}
	// prev is now head of input list
	return prev
}
