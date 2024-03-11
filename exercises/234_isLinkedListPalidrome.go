package exercises

import "fmt"

// Given the head of a singly linked list, return true if it is a palindrome or false otherwise.

func IsLinkedListPalindrome(head *ListNode) bool {
	// convert to slice
	convertedSlice := []int{}

	for head != nil {
		convertedSlice = append(convertedSlice, head.Val)
		head = head.Next
	}

	for i := range convertedSlice {
		if convertedSlice[len(convertedSlice)-1-i] != convertedSlice[i] {
			return false
		}
	}

	return true
}

func IsLinkedListPalindromeTester() bool {
	list1 := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}

	fmt.Print(ReverseLinkedList(list1)) // true

	return true
}
