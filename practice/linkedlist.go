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

func newList(value interface{}) linkedList {
	ll := linkedList{
		head:   newNode(value),
		length: 1,
	}
	ll.tail = ll.head
	return ll
}

func newNode(value interface{}) *node {
	return &node{
		next:  nil,
		prev:  nil,
		value: value,
	}
}

func (ll *linkedList) prepend(value interface{}) {
	newNode := newNode(value)
	newNode.next = ll.head
	ll.head.prev = newNode
	ll.head = newNode
	ll.length++

}
func (ll *linkedList) append(value interface{}) {
	newNode := newNode(value)
	newNode.prev = ll.tail
	ll.tail.next = newNode
	ll.tail = newNode
	ll.length++
}

func (ll *linkedList) print() {
	head := ll.head
	for head != nil {
		fmt.Println(head.value, "--> ", "prev: ", head.prev, "next: ", head.next)
		head = head.next
	}
}

func (ll *linkedList) traverse(index int) *node {
	node := ll.head
	if index >= ll.length {
		index = ll.length - 1
	}
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
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
	previousNode.next.prev = newNode
	previousNode.next = newNode
	ll.length++
}

func (ll *linkedList) remove(index int) {
	previousNode := ll.traverse(index - 1)
	previousNode.next.next.prev = previousNode
	previousNode.next = previousNode.next.next
	//previousNode.next.prev = previousNode
	ll.length--
}

func (ll *linkedList) reverse() {
	if ll.length > 1 {
		first := ll.head
		second := first.next

		for second != nil {
			temp := second.next
			second.prev = second.next
			second.next = first
			first = second
			second = temp
		}
		ll.head.next = nil
		ll.head = first

	}
}

func main() {
	myList := newList(20)
	myList.prepend(5)
	myList.append(30)
	myList.insert(1, 10)
	myList.insert(2, 15)
	myList.insert(4, 25)
	//	myList.insert(3, 55)
	//myList.insert(0, 0)
	// myList.insert(5, 16)
	// myList.insert(6, 17)
	//myList.remove(3)
	myList.print()

	fmt.Println("reverse")
	myList.reverse()
	myList.print()

	str1 := "c"
	str2 := "b"

	fmt.Println(str1 > str2)
}
