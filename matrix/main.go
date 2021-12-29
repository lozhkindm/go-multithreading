package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	MatrixSize = 250
)

var (
	matrixA = [MatrixSize][MatrixSize]int{
		{3, 1, -4},
		{2, -3, 1},
		{5, -2, 0},
	}
	matrixB = [MatrixSize][MatrixSize]int{
		{1, -2, -1},
		{0, 5, 4},
		{-1, -2, 3},
	}
	result = [MatrixSize][MatrixSize]int{}
	rwLock = sync.RWMutex{}
	cond   = sync.NewCond(rwLock.RLocker())
	wg     = sync.WaitGroup{}
)

func populateMatrix(matrix *[MatrixSize][MatrixSize]int) {
	for row := 0; row < MatrixSize; row++ {
		for col := 0; col < MatrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}

func calcRow(row int) {
	rwLock.RLock()
	for {
		wg.Done()
		cond.Wait()
		for col := 0; col < MatrixSize; col++ {
			for i := 0; i < MatrixSize; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][col]
			}
		}
	}
}

func main() {
	wg.Add(MatrixSize)
	for row := 0; row < MatrixSize; row++ {
		go calcRow(row)
	}

	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Wait()
		rwLock.Lock()
		populateMatrix(&matrixA)
		populateMatrix(&matrixB)
		wg.Add(MatrixSize)
		rwLock.Unlock()
		cond.Broadcast()
	}
	elapsed := time.Since(start)
	fmt.Printf("Processing took: %v", elapsed)
}
