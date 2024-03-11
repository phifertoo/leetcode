package exercises

import "fmt"

// Given the head of a singly linked list, reverse the list, and return the reversed list.

func ReverseLinkedList(head *ListNode) *ListNode {
	var previous *ListNode
	current := head
	for current != nil {
		// we need to adjust the "Next" to the previous node
		next := current.Next    // Store next node
		current.Next = previous // Reverse current node's pointer
		previous = current      // Move pointers one step forward
		current = next
	}
	return previous // New head of the reversed list
}

func ReverseLinkedListTester() bool {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}

	fmt.Print(ReverseLinkedList(list)) // [1, 1, 2, 3, 4, 4]

	return true
}
