package main

import "fmt"

func main() {
	singlyLinkedList := NewSinglyLinkedList(10)
	singlyLinkedList.append(15)
	singlyLinkedList.prepend(5)
	singlyLinkedList.append(20)
	singlyLinkedList.prepend(1)
	singlyLinkedList.insert(2, 100)
	singlyLinkedList.remove(4)
	singlyLinkedList.print()

	fmt.Println("Doubly linked list...")
	doublyLinkedList := newDoublylinkedList(10)
	//doublyLinkedList.print()
	doublyLinkedList.prepend(2)
	doublyLinkedList.append(3)
	doublyLinkedList.prepend(1)
	doublyLinkedList.append(5)
	doublyLinkedList.print()

	fmt.Println("removing at index 2")
	doublyLinkedList.remove(2)
	doublyLinkedList.print()

	fmt.Println("inserting 100 at index 3")
	doublyLinkedList.insert(3, 4)
	doublyLinkedList.print()

	fmt.Println("inserting 999 at index 5")
	doublyLinkedList.append(6)
	doublyLinkedList.print()

	fmt.Println("reversing the linked list")
	doublyLinkedList.reverse()
	doublyLinkedList.print()

	fmt.Println("reversing the linked list to reset")
	doublyLinkedList.reverse()
	doublyLinkedList.print()

	doublyLinkedList.prepend(0)
	doublyLinkedList.append(10)
	doublyLinkedList.print()
	fmt.Println("=====")

	doublyLinkedList.insert(7, 7)
	doublyLinkedList.insert(8, 8)
	doublyLinkedList.insert(9, 9)
	doublyLinkedList.print()

	fmt.Println("segmented reverse")
	doublyLinkedList.head = doublyLinkedList.reverseSegmented(5, doublyLinkedList.head)
	//doublyLinkedList.head.next = nil
	doublyLinkedList.print()
}
