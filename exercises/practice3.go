package exercises

import "fmt"

// type Element struct {
// 	Value     int
// 	Frequency int
// }

// func Practice3(nums []int, k int) []int {
// 	// make frequency map
// 	freqMap := make(map[int]int)
// 	for _, num := range nums {
// 		freqMap[num]++
// 	}

// 	// iterate over frequency map to add each element to the heap
// 	// if the heap is
// 	heap := []Element{}
// 	for value, frequency := range freqMap {
// 		// add element to end
// 		heap = append(heap, Element{
// 			Value:     value,
// 			Frequency: frequency,
// 		})
// 		// move element up the tree
// 		index := len(heap) - 1
// 		for index > 0 {
// 			parentIndex := (index - 1) / 2 // Calculate parent index
// 			if heap[index].Frequency >= heap[parentIndex].Frequency {
// 				break // Stop if the current element is not less than its parent
// 			}
// 			// Swap the current element with its parent
// 			heap[index], heap[parentIndex] = heap[parentIndex], heap[index]
// 			index = parentIndex // Move up to the parent index
// 		}

// 		// if the heap is too big, remove the first (smallest)
// 		if len(heap) > k {
// 			min := heap[0]
// 			heap[0] = heap[len(heap)-1]
// 			heap = heap[:len(heap)-1]

// 			//
// 			lastIndex := len(heap) - 1
// 			downIndex := 0
// 			for {
// 				leftChildIndex := 2*downIndex + 1
// 				rightChildIndex := 2*downIndex + 2

// 				smallestIndex := downIndex // Assume the current index is the smallest

// 				// Find the smallest among the current, left child, and right child
// 				if leftChildIndex <= lastIndex && heap[leftChildIndex].Frequency < heap[smallestIndex].Frequency {
// 					smallestIndex = leftChildIndex
// 				}

// 				if rightChildIndex <= lastIndex && heap[rightChildIndex].Frequency < heap[smallestIndex].Frequency {
// 					smallestIndex = rightChildIndex
// 				}

// 				if smallestIndex == downIndex {
// 					break // If the smallest is the current element, the heap property is satisfied
// 				}

// 				// Swap the current element with the smallest of its children
// 				heap[downIndex], heap[smallestIndex] = heap[smallestIndex], heap[downIndex]
// 				downIndex = smallestIndex // Move down to the smallest child's index
// 			}
// 		}
// 	}

// 	// iterate over frequency map
// 	// result := make([]int, k)
// 	// for i := k - 1; i >= 0; i-- {
// 	// 	result[k - i - 1] = heap.Pop(h).(Element).Value
// 	// }
// 	// return result
// }

func Practice3Tester() bool {
	fmt.Print("no")

	return true
}
