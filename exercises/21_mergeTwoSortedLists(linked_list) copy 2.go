package exercises

import "fmt"

// You are given the heads of two sorted linked lists list1 and list2.

// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

// Return the head of the merged linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

// The linked list is ordered  by ListNodes.
// Each ListNode has 1 number and the "next" ListNode
// 1. make a "head" which represents the sorted list
// 2. make a pointer to the head called "current" to append each ListNode
// 3. iterate through both linkedLists.
// 4. if ListNode1.Value <= ListNode2.Value => append ListNode1, append the remaining to current.Next
// 5. append the entire
func mergeTwoLists(listNode1 *ListNode, listNode2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head

	for listNode1 != nil && listNode2 != nil {
		if listNode1.Val <= listNode2.Val {
			// store the next value in current.Next
			current.Next = listNode1
			// move listNode1 forward
			listNode1 = listNode1.Next
		} else {
			// store the next value in current.Next
			current.Next = listNode2
			// move listNode1 forward
			listNode2 = listNode2.Next
		}
		// set the current to the value we stored in current.Next
		current = current.Next
	}

	if listNode1 == nil {
		current.Next = listNode2
	}
	if listNode2 == nil {
		current.Next = listNode1
	}

	return head.Next
}

func MergeTwoListsTester() bool {
	node3 := &ListNode{Val: 4, Next: nil}   // Last node, points to nil
	node2 := &ListNode{Val: 2, Next: node3} // Middle node, points to last node
	node1 := &ListNode{Val: 1, Next: node2} // First node, points to middle node

	nodeC := &ListNode{Val: 4, Next: nil}   // Last node, points to nil
	nodeB := &ListNode{Val: 3, Next: nodeC} // Middle node, points to last node
	nodeA := &ListNode{Val: 1, Next: nodeB} // First node, points to middle node

	node0 := &ListNode{Val: 0, Next: nil} // First node, points to middle node

	fmt.Print(mergeTwoLists(node1, nodeA)) // [1, 1, 2, 3, 4, 4]
	fmt.Print(mergeTwoLists(nil, nil))     //[]
	fmt.Print(mergeTwoLists(nil, node0))   //[0]

	return true
}
