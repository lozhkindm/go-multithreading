package main

import (
	"fmt"
	"time"
)

func waitOnBarrier(name string, duration time.Duration, barrier *Barrier) {
	for {
		fmt.Println(name, "is running")
		time.Sleep(duration)
		fmt.Println(name, "is waiting on a barrier")
		barrier.Wait()
	}
}

func main() {
	barrier := NewBarrier(2)
	go waitOnBarrier("red robot", 4*time.Second, barrier)
	go waitOnBarrier("blue robot", 10*time.Second, barrier)
	time.Sleep(100 * time.Second)
}
