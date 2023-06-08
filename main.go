package main

import "math"

type Kubus struct {
	sisi float64
}

func (k Kubus) keliling() float64 {
	return k.sisi * 12
}

func (k Kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k Kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}
