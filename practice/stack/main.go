package main

import "fmt"

func main() {
	myStack := newStack(100)
	i := 10
	for i >= 1 {
		myStack.push(i)
		i--
	}

	for i < 5 {
		myStack.pop()
		i++
	}
	fmt.Println("peek result :", myStack.peek().value)
	myStack.print()
	myStack.traverse(4)

	fmt.Println("array stack")

	as := newArrayStack(2)
	for i := 0; i < 10; i++ {
		as.push(i * 4)
	}
	as.pop()
	as.pop()
	as.pop()
	as.print()
	fmt.Println("peek value: ", as.peek())
	as.traverse(5)

}
