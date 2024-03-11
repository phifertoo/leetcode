package exercises

import "fmt"

// You are given the heads of two sorted linked lists list1 and list2.

// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

// Return the head of the merged linked list.

func hasCycle(head *ListNode) bool {
	visited := make(map[ListNode]bool)
	for head != nil {
		if _, exists := visited[*head]; exists {
			return true // Cycle detected
		}
		visited[*head] = true
		head = head.Next
	}
	return false // No cycle
}

func HasCycleTester() bool {
	node4 := &ListNode{Val: 4, Next: nil}   // Last node, points to nil
	node3 := &ListNode{Val: 0, Next: node4} // Last node, points to nil
	node2 := &ListNode{Val: 2, Next: node3} // Middle node, points to last node
	node1 := &ListNode{Val: 3, Next: node2} // First node, points to middle node
	node4.Next = node2

	nodeB := &ListNode{Val: 2, Next: nil}   // Middle node, points to last node
	nodeA := &ListNode{Val: 1, Next: nodeB} // First node, points to middle node
	nodeB.Next = nodeA

	node0 := &ListNode{Val: 1, Next: nil} // First node, points to middle node

	fmt.Print(hasCycle(node1)) // [1, 1, 2, 3, 4, 4]
	fmt.Print(hasCycle(nodeA)) //[]
	fmt.Print(hasCycle(node0)) //[0]

	return true
}
