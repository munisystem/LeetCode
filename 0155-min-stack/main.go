package main

type MinStack struct {
	min   int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: []int{}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.stack) == 1 || this.min > val {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	elm := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if this.min != elm {
		return
	}
	for i := 0; i < len(this.stack); i++ {
		if i == 0 || this.min > this.stack[i] {
			this.min = this.stack[i]
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
