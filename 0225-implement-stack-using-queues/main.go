package main

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{queue: []int{}}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
	for i := 0; i < len(this.queue)-1; i++ {
		tmp := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, tmp)
	}
}

func (this *MyStack) Pop() int {
	v := this.queue[0]
	this.queue = this.queue[1:]
	return v
}

func (this *MyStack) Top() int {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
