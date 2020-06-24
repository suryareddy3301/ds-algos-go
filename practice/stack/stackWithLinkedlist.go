package main

import "fmt"

//push pop peek

type node struct {
	value interface{}
	next  *node
}

type linkedListStack struct {
	top    *node
	length int
}

func newNode(value interface{}) *node {
	return &node{
		value: value,
	}
}

func newLinkedListStack(value interface{}) linkedListStack {
	node := newNode(value)
	stack := linkedListStack{
		top:    node,
		length: 1,
	}
	return stack
}

func (s *linkedListStack) push(value interface{}) {
	newNode := newNode(value)
	newNode.next = s.top
	s.top = newNode
}

func (s *linkedListStack) pop() {
	s.top = s.top.next
}

func (s *linkedListStack) peek() {
	fmt.Println("peek:", s.top.value)
}

func (s *linkedListStack) print() {
	top := s.top
	for top != nil {
		fmt.Println(top.value)
		top = top.next
	}
}
