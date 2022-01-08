package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

const (
	TotalAccounts  = 50000
	MaxAmountTrans = 10
	InitAmount     = 100
	Threads        = 4
)

func performTransactions(ledger *[TotalAccounts]int32, locks *[TotalAccounts]sync.Locker, totalTrans *int64) {
	for {
		accountFrom := rand.Intn(TotalAccounts)
		accountTo := rand.Intn(TotalAccounts)
		for accountFrom == accountTo {
			accountTo = rand.Intn(TotalAccounts)
		}
		amountTrans := rand.Int31n(MaxAmountTrans)
		toLock := []int{accountFrom, accountTo}
		sort.Ints(toLock)

		locks[toLock[0]].Lock()
		locks[toLock[1]].Lock()
		atomic.AddInt32(&ledger[accountFrom], -amountTrans)
		atomic.AddInt32(&ledger[accountTo], amountTrans)
		atomic.AddInt64(totalTrans, 1)
		locks[toLock[0]].Unlock()
		locks[toLock[1]].Unlock()
	}
}

func main() {
	fmt.Printf("total accounts: %d, total threads using a Spinlock: %d\n", TotalAccounts, Threads)
	var ledger [TotalAccounts]int32
	var locks [TotalAccounts]sync.Locker
	var totalTrans int64
	for i := 0; i < TotalAccounts; i++ {
		ledger[i] = InitAmount
		locks[i] = NewSpinlock()
	}
	for i := 0; i < Threads; i++ {
		go performTransactions(&ledger, &locks, &totalTrans)
	}
	for {
		time.Sleep(2000 * time.Millisecond)
		var sum int32
		for i := 0; i < TotalAccounts; i++ {
			locks[i].Lock()
		}
		for i := 0; i < TotalAccounts; i++ {
			sum += ledger[i]
		}
		for i := 0; i < TotalAccounts; i++ {
			locks[i].Unlock()
		}
		fmt.Printf("Total sum: %d, total transactions: %d\n", sum, totalTrans)
	}
}
