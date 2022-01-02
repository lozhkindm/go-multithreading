package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock1 = sync.Mutex{}
	lock2 = sync.Mutex{}
)

func blueRobot() {
	for {
		fmt.Println("Blue: acquiring lock1")
		lock1.Lock()
		fmt.Println("Blue: acquiring lock2")
		lock2.Lock()
		fmt.Println("Blue: locks acquired")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Blue: locks released")
	}
}

func redRobot() {
	for {
		fmt.Println("Red: acquiring lock2")
		lock2.Lock()
		fmt.Println("Red: acquiring lock1")
		lock1.Lock()
		fmt.Println("Red: locks acquired")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Red: locks released")
	}
}

func main() {
	go blueRobot()
	go redRobot()
	time.Sleep(20 * time.Second)
	fmt.Println("Done")
}
