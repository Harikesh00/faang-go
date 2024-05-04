package main

// 27
type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.min) == 0 || val <= this.GetMin() {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	if this.stack[len(this.stack)-1] == this.GetMin() {
		this.min = this.min[:len(this.min)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

// func main() {
//     obj := Constructor()
//     obj.Push(-2)
//     obj.Push(0)
//     obj.Push(-3)
//     fmt.Println(obj.GetMin()) // Output: -3
//     obj.Pop()
//     fmt.Println(obj.Top())    // Output: 0
//     fmt.Println(obj.GetMin()) // Output: -2
// }
