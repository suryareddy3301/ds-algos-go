package main

import "fmt"

type node struct {
	value interface{}
	next  *node
}

type stack struct {
	top *node
}

func newNode(value interface{}) *node {
	return &node{
		value: value,
		next:  nil,
	}
}

func newStack(value interface{}) stack {
	return stack{
		top: newNode(value),
	}
}

func (s *stack) push(value interface{}) {
	newNode := newNode(value)
	newNode.next = s.top
	s.top = newNode
}

func (s *stack) pop() {
	s.top = s.top.next
}

func (s *stack) print() {
	node := s.top
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

func (s *stack) peek() *node {
	return s.top
}

func (s *stack) traverse(index int) {
	current := s.top
	for i := 0; i < index; i++ {
		current = current.next
	}
	fmt.Println("value at index ", index, " is: ", current.value)
}
