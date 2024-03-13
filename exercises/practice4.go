package exercises

import "fmt"

// parentIndex = (index - 1) / 2
// leftChild (index * 2) + 1
// rightChild (index * 2) + 2

// Element is a struct that holds an integer value and its frequency in the given slice of integers.
type Element struct {
	Value     int
	Frequency int
}

// Minheap is a struct that represents a min heap of Elements. It has methods to insert elements, remove the minimum element, and to maintain the heap property (through upHeap and downHeap operations).
type MinHeap struct {
	Elements []Element
}

// Adds a new Element to the heap, appends it to the Elements slice, and then adjusts the heap to maintain the min heap property using the upHeap method.
func (h *MinHeap) insertAndSort(item Element) {
	h.Elements = append(h.Elements, item) // Add the new element to the end
	// after you insert, you need to make sure the child frequency > parent frequency
	currentIndex := len(h.Elements) - 1 // current index is set to the element we just appended
	for currentIndex > 0 {
		parentIndex := (currentIndex - 1) / 2 // Calculate parent index
		if h.Elements[currentIndex].Frequency >= h.Elements[parentIndex].Frequency {
			break // Stop if the current element is greater than its parent
		}
		// Swap the current element with its parent
		h.Elements[currentIndex], h.Elements[parentIndex] = h.Elements[parentIndex], h.Elements[currentIndex]
		currentIndex = parentIndex // Move up to the parent index
	}
}

// Removes and returns the root element (the minimum element) of the heap. It replaces the root with the last element in the heap, removes the last element from the slice, and then adjusts the heap to maintain the min heap property using the downHeap method.
func (h *MinHeap) removeMin() Element {
	if len(h.Elements) == 0 {
		return Element{} // Return an empty Element if the heap is empty
	}
	min := h.Elements[0]                          // The min element is always at the root
	h.Elements[0] = h.Elements[len(h.Elements)-1] // Replace the root with the last element
	h.Elements = h.Elements[:len(h.Elements)-1]   // Remove the last element
	h.downHeap(0)                                 // moves the root to the correct position
	return min
}

// Ensures the min heap property is maintained after the root element is removed. It compares the new root with its children and swaps it with the smaller child. This process is repeated until the element is in the correct position.

func (h *MinHeap) downHeap(index int) {
	// index is the root

	lastIndex := len(h.Elements) - 1
	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2

		smallestIndex := index // Assume the current index is the smallest

		// Find the smallest among the current, left child, and right child
		if leftChildIndex <= lastIndex && h.Elements[leftChildIndex].Frequency < h.Elements[smallestIndex].Frequency {
			smallestIndex = leftChildIndex
		}

		if rightChildIndex <= lastIndex && h.Elements[rightChildIndex].Frequency < h.Elements[smallestIndex].Frequency {
			smallestIndex = rightChildIndex
		}

		if smallestIndex == index {
			break // If the smallest is the current element, the heap property is satisfied
		}

		// Swap the current element with the smallest of its children
		h.Elements[index], h.Elements[smallestIndex] = h.Elements[smallestIndex], h.Elements[index]
		index = smallestIndex // Move down to the smallest child's index
	}
}

// topKFrequent finds the k most frequent elements in an array.
func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++ // Count the frequency of each element
	}

	h := MinHeap{}

	for value, frequency := range freqMap {
		element := Element{Value: value, Frequency: frequency}
		h.insertAndSort(element) // Insert each element with its frequency into the heap
		if len(h.Elements) > k {
			h.removeMin() // Keep only the k most frequent elements in the heap
		}
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[k-i-1] = h.removeMin().Value // Extract elements from the heap, in reverse order
	}
	return result
}

func Practice4Tester() bool {
	fmt.Print(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))

	return true
}

// Ensures the min heap property is maintained after insertion. It compares the inserted element with its parent; if the parent's frequency is greater, they are swapped. This process is repeated until the element is in the correct position.
// element to the left of the inserted element
// func (h *MinHeap) upHeap(index int) {
// 	for index > 0 {
// 		parentIndex := (index - 1) / 2 // Calculate parent index
// 		if h.Elements[index].Frequency >= h.Elements[parentIndex].Frequency {
// 			break // Stop if the current element is greater than its parent
// 		}
// 		// Swap the current element with its parent
// 		h.Elements[index], h.Elements[parentIndex] = h.Elements[parentIndex], h.Elements[index]
// 		index = parentIndex // Move up to the parent index
// 	}
// }
