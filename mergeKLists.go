package main

import (
	"container/heap"
)

// 49
// ListNode represents a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// ListNodeHeap is a min-heap (priority queue) of ListNode pointers
type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int            { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *ListNodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	// Create a min-heap and push the head nodes of all lists into it
	minHeap := &ListNodeHeap{}
	heap.Init(minHeap)
	for _, node := range lists {
		if node != nil {
			heap.Push(minHeap, node)
		}
	}

	// Create a dummy node to serve as the head of the merged list
	dummy := &ListNode{}
	current := dummy

	// Merge the lists by extracting the smallest node from the min-heap
	for minHeap.Len() > 0 {
		// Extract the smallest node from the min-heap
		smallest := heap.Pop(minHeap).(*ListNode)
		current.Next = smallest
		current = current.Next

		// Push the next node from the same list into the min-heap
		if smallest.Next != nil {
			heap.Push(minHeap, smallest.Next)
		}
	}

	return dummy.Next
}

// func main() {
// 	// Example usage:
// 	// Create linked lists: 1->4->5, 1->3->4, 2->6
// 	list1 := &ListNode{Val: 1}
// 	list1.Next = &ListNode{Val: 4}
// 	list1.Next.Next = &ListNode{Val: 5}

// 	list2 := &ListNode{Val: 1}
// 	list2.Next = &ListNode{Val: 3}
// 	list2.Next.Next = &ListNode{Val: 4}

// 	list3 := &ListNode{Val: 2}
// 	list3.Next = &ListNode{Val: 6}

// 	// Merge the lists
// 	merged := mergeKLists([]*ListNode{list1, list2, list3})

// 	// Print the merged list
// 	for merged != nil {
// 		fmt.Printf("%d ", merged.Val)
// 		merged = merged.Next
// 	}
// 	// Output: 1 1 2 3 4 4 5 6
// }
