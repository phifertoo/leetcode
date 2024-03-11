package exercises

import "fmt"

// Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. If the two linked lists have no intersection at all, return null.

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA := headA
	pB := headB

	// Continue traversing until both pointers meet or both reach the end (null)
	for pA != pB {
		// If pA reaches the end of list A, switch to the head of list B
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		// If pB reaches the end of list B, switch to the head of list A
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	// pA is either the intersection node or null if there is no intersection
	return pA
}

func GetIntersectionNodeTester() bool {
	intersection := &ListNode{Val: 2, Next: &ListNode{Val: 4}}

	// Create list A
	headA := &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: intersection}}}

	// Create list B
	headB := &ListNode{Val: 3, Next: intersection}

	fmt.Print(getIntersectionNode(headA, headB)) // true

	return true
}
