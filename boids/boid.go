package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func (b *Boid) moveOne() {
	accel := b.calcAcceleration()
	lock.Lock()
	b.velocity = b.velocity.Add(accel).limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id

	next := b.position.Add(b.velocity)
	if next.x >= screenWidth || next.x < 0 {
		b.velocity = Vector2D{
			x: -b.velocity.x,
			y: b.velocity.y,
		}
	}
	if next.y >= screenHeight || next.y < 0 {
		b.velocity = Vector2D{
			x: b.velocity.x,
			y: -b.velocity.y,
		}
	}
	lock.Unlock()
}

func (b *Boid) calcAcceleration() Vector2D {
	up, low := b.position.AddVal(viewRadius), b.position.AddVal(-viewRadius)
	avgVelocity := Vector2D{x: 0, y: 0}
	avgPosition := Vector2D{x: 0, y: 0}
	count := 0.0

	lock.Lock()
	for i := math.Max(low.x, 0); i <= math.Min(up.x, screenWidth); i++ {
		for j := math.Max(low.y, 0); j <= math.Min(up.y, screenHeight); j++ {
			if bid := boidMap[int(i)][int(j)]; bid != -1 && bid != b.id {
				if dist := boids[bid].position.Distance(b.position); dist < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(boids[bid].velocity)
					avgPosition = avgPosition.Add(boids[bid].position)
				}
			}
		}
	}
	lock.Unlock()

	accel := Vector2D{x: 0, y: 0}
	if count > 0 {
		avgVelocity = avgVelocity.DivVal(count)
		avgPosition = avgPosition.DivVal(count)
		accelAlignment := avgVelocity.Sub(b.velocity).MultiVal(adjRate)
		accelCohesion := avgPosition.Sub(b.position).MultiVal(adjRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion)
	}

	return accel
}

func createBoid(bid int) {
	b := Boid{
		position: Vector2D{
			x: rand.Float64() * screenWidth,
			y: rand.Float64() * screenHeight,
		},
		velocity: Vector2D{
			x: (rand.Float64() * 2) - 1.0,
			y: (rand.Float64() * 2) - 1.0,
		},
		id: bid,
	}
	boidMap[int(b.position.x)][int(b.position.y)] = bid
	boids[bid] = &b
	go b.start()
}
