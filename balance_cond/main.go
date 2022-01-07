package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	money          = 100
	lock           = sync.Mutex{}
	monedDeposited = sync.NewCond(&lock)
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10
		fmt.Println("Stingy sees balance:", money)
		monedDeposited.Signal()
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Stingy Done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		for money-20 < 0 {
			monedDeposited.Wait()
		}
		money -= 20
		fmt.Println("Spendy sees balance:", money)
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Spendy Done")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println(money)
}
