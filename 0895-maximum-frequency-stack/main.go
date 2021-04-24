package main

import (
	"math"
)

type FreqStack struct {
	stack   map[int][]int
	freq    map[int]int
	maxFreq int
}

func Constructor() FreqStack {
	return FreqStack{
		stack:   map[int][]int{},
		freq:    map[int]int{},
		maxFreq: 0,
	}
}

func (this *FreqStack) Push(x int) {
	this.freq[x]++
	this.maxFreq = int(math.Max(float64(this.maxFreq), float64(this.freq[x])))
	this.stack[this.freq[x]] = append(this.stack[this.freq[x]], x)
}

func (this *FreqStack) Pop() int {
	t := this.stack[this.maxFreq][len(this.stack[this.maxFreq])-1]
	this.stack[this.maxFreq] = this.stack[this.maxFreq][:len(this.stack[this.maxFreq])-1]
	this.freq[t]--
	if len(this.stack[this.maxFreq]) == 0 {
		this.maxFreq--
	}
	return t
}
