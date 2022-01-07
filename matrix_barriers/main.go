package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MatrixSize = 250
)

var (
	matrixA   = [MatrixSize][MatrixSize]int{}
	matrixB   = [MatrixSize][MatrixSize]int{}
	result    = [MatrixSize][MatrixSize]int{}
	calcStart = NewBarrier(MatrixSize + 1)
	calcEnd   = NewBarrier(MatrixSize + 1)
)

func populateMatrix(matrix *[MatrixSize][MatrixSize]int) {
	for row := 0; row < MatrixSize; row++ {
		for col := 0; col < MatrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}

func calcRow(row int) {
	for {
		calcStart.Wait()
		for col := 0; col < MatrixSize; col++ {
			for i := 0; i < MatrixSize; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][col]
			}
		}
		calcEnd.Wait()
	}
}

func main() {
	for row := 0; row < MatrixSize; row++ {
		go calcRow(row)
	}

	start := time.Now()
	for i := 0; i < 100; i++ {
		populateMatrix(&matrixA)
		populateMatrix(&matrixB)
		calcStart.Wait()
		calcEnd.Wait()
	}
	elapsed := time.Since(start)
	fmt.Printf("Processing took: %v", elapsed)
}
