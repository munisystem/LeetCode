package main

import (
	"math"
	"math/rand"
)

type Solution struct {
	radius float64
	x      float64
	y      float64
	rand   *rand.Rand
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	r := rand.New(rand.NewSource(99))
	return Solution{radius: radius, x: x_center, y: y_center, rand: r}
}

func (this *Solution) RandPoint() []float64 {
	l := math.Sqrt(this.rand.Float64()) * this.radius
	d := this.rand.Float64() * 2 * math.Pi
	x := this.x + l*math.Cos(d)
	y := this.y + l*math.Sin(d)
	return []float64{x, y}
}
