package fox

// define Element struct
type Element struct {
	Value     int
	Frequency int
}

// define heap
type MinHeap struct {
	Elements []Element
}

func (h *MinHeap) insertAndSort(element Element) {
	h.Elements = append(h.Elements, element)
	currentIndex := len(h.Elements) - 1
	for currentIndex > 0 {
		parentIndex := (currentIndex - 1) / 2
		if h.Elements[parentIndex].Frequency < h.Elements[currentIndex].Frequency {
			break
		}
		h.Elements[currentIndex], h.Elements[parentIndex] = h.Elements[parentIndex], h.Elements[currentIndex]
		currentIndex = parentIndex
	}
}

func (h *MinHeap) removeMin() Element {
	// assume the minimum is the root
	min := h.Elements[0]
	h.Elements[0] = h.Elements[len(h.Elements)-1]
	h.Elements = h.Elements[:len(h.Elements)-1]
	// now the root is out of order so we need to move it down
	h.downHeap(0)
	return min
}

func (h *MinHeap) downHeap(index int) {
	lastIndex := len(h.Elements) - 1
	for {
		leftChildIndex := index*2 + 1
		rightChildIndex := index*2 + 2

		// assume smallest index is current index then check
		smallestIndex := index
		if leftChildIndex <= lastIndex && h.Elements[leftChildIndex].Frequency < h.Elements[smallestIndex].Frequency {
			smallestIndex = leftChildIndex
		}
		if rightChildIndex <= lastIndex && h.Elements[rightChildIndex].Frequency < h.Elements[smallestIndex].Frequency {
			smallestIndex = rightChildIndex
		}
		// keep looping until the smallest index  = index
		if smallestIndex == index {
			break
		}
		// swap smallest index with the current index
		h.Elements[smallestIndex], h.Elements[index] = h.Elements[index], h.Elements[smallestIndex]
		index = smallestIndex
	}
}

func (h *MinHeap) HighestFreq(nums []int, k int) []int {
	// make frequency map
	freqMap := make(map[int]int)
	for _, num := range nums {
		_, exists := freqMap[num]
		if exists {
			freqMap[num]++
		} else {
			freqMap[num] = 1
		}
	}

	// fmt.Printf("start:%+v:end\n", freqMap)

	// itereate through frequency map to append element to heap
	for value, frequency := range freqMap {
		// make element
		element := Element{
			Value:     value,
			Frequency: frequency,
		}
		h.insertAndSort(element)
		// if the length of the heap is longer than k, you need to swap the root with the new element, then cut off the root (last element)
		if len(h.Elements) > k {
			h.removeMin()
		}
	}

	// // return the array with the minimum k elements
	// output := make([]int, k)
	// for i := 0; i < k; i++ {
	// 	// append to output in reverse order
	// 	output[k-i-1] = h.removeMin().Value // Extract elements from the heap, in reverse order
	// }
	output := []int{}
	for _, element := range h.Elements {
		output = append(output, element.Value)
	}

	return output
}
