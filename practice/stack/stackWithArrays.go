package main

import "fmt"

type arrayStack struct {
	content []interface{}
}

func newArrayStack(value interface{}) arrayStack {
	content := make([]interface{}, 0)
	content = append(content, value)

	return arrayStack{
		content: content,
	}

}

func (as *arrayStack) push(value interface{}) {
	as.content = append(as.content, value)
}

func (as *arrayStack) pop() {
	as.content = as.content[:len(as.content)-1]
}

func (as *arrayStack) peek() interface{} {
	return as.content[len(as.content)-1]
}

func (as *arrayStack) print() {
	len := len(as.content) - 1
	for i := len; i >= 0; i-- {
		fmt.Println(as.content[i])
	}
}

func (as *arrayStack) traverse(index int) {
	var val interface{}
	counter := 0
	for i := len(as.content) - 1; i >= 0; i-- {
		if counter == index {
			val = as.content[i]
			break
		}
		counter++
	}
	fmt.Println("value at index ", index, " is: ", val)
}
