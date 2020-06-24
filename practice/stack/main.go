package main

import "fmt"

func main() {
	linkedListStack := newLinkedListStack("apple")
	linkedListStack.push("ball")
	linkedListStack.push("cat")
	linkedListStack.print()
	linkedListStack.peek()
	fmt.Println("popping..")
	linkedListStack.pop()
	linkedListStack.print()
	fmt.Println("popping...")
	linkedListStack.pop()
	linkedListStack.print()
}
