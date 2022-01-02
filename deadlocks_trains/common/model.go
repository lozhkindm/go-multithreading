package common

import "sync"

type Train struct {
	Id     int
	Length int
	Front  int
}

type Intersection struct {
	Id       int
	Mx       sync.Mutex
	LockedBy int
}

type Crossing struct {
	Position     int
	Intersection *Intersection
}
