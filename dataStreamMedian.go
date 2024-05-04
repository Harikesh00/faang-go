package main

// 45
/*
type MaxHeap []int
type MinHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	maxHeap *MaxHeap
	minHeap *MinHeap
}

func Constructor() MedianFinder {
	maxHeap := &MaxHeap{}
	minHeap := &MinHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)
	return MedianFinder{maxHeap, minHeap}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.maxHeap.Len() == 0 || num <= (*mf.maxHeap)[0] {
		heap.Push(mf.maxHeap, num)
	} else {
		heap.Push(mf.minHeap, num)
	}

	// Balance the heaps
	if mf.maxHeap.Len()-mf.minHeap.Len() > 1 {
		heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))
	} else if mf.minHeap.Len()-mf.maxHeap.Len() > 1 {
		heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.maxHeap.Len() == mf.minHeap.Len() {
		return float64((*mf.maxHeap)[0]+(*mf.minHeap)[0]) / 2
	} else if mf.maxHeap.Len() > mf.minHeap.Len() {
		return float64((*mf.maxHeap)[0])
	} else {
		return float64((*mf.minHeap)[0])
	}
}

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian()) // Output: 1.5
	obj.AddNum(3)
	fmt.Println(obj.FindMedian()) // Output: 2
}
*/
