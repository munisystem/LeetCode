package main

type OrderedStream struct {
	chunks  []string
	pointer int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		chunks:  make([]string, n, n),
		pointer: 0,
	}
}

func (this *OrderedStream) Insert(id int, value string) []string {
	this.chunks[id-1] = value
	start := this.pointer
	for this.pointer < len(this.chunks) {
		if this.chunks[this.pointer] == "" {
			break
		}
		this.pointer++
	}
	return this.chunks[start:this.pointer]
}
