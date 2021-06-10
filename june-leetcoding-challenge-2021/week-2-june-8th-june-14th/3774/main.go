package main

type MyCalendar struct {
	books [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{books: [][]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for i := 0; i < len(this.books); i++ {
		book := this.books[i]
		if book[0] < end && start < book[1] {
			return false
		}
	}
	this.books = append(this.books, append([]int{}, start, end))
	return true
}
