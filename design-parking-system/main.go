package main

type car int

const (
	b car = iota + 1
	m
	s
)

type ParkingSystem struct {
	parking map[car]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	parking := map[car]int{}
	parking[s] = small
	parking[m] = medium
	parking[b] = big
	return ParkingSystem{parking: parking}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.parking[car(carType)] == 0 {
		return false
	}
	this.parking[car(carType)]--
	return true
}
