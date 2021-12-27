package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const NumberOfThreads = 2

var (
	r  = regexp.MustCompile(`\((\d+),(\d+)\)`)
	wg = sync.WaitGroup{}
)

type Point2D struct {
	x int
	y int
}

func findArea(inputChan chan string) {
	for pts := range inputChan {
		var points []Point2D

		r.FindAllStringSubmatch(pts, -1)
		for _, p := range r.FindAllStringSubmatch(pts, -1) {
			x, _ := strconv.Atoi(p[1])
			y, _ := strconv.Atoi(p[2])
			points = append(points, Point2D{x: x, y: y})
		}

		area := 0.0
		for i := 0; i < len(points); i++ {
			p1, p2 := points[i], points[(i+1)%len(points)]
			area += float64(p1.x*p2.y) - float64(p1.y*p2.x)
		}
		fmt.Printf("%v\n", math.Abs(area)/2)
	}
	wg.Done()
}

func main() {
	path, _ := filepath.Abs("./shoelace")
	data, _ := ioutil.ReadFile(filepath.Join(path, "polygons.txt"))
	pgons := string(data)

	inputChan := make(chan string, 1000)
	for i := 0; i < NumberOfThreads; i++ {
		go findArea(inputChan)
	}
	wg.Add(NumberOfThreads)

	start := time.Now()
	for _, line := range strings.Split(pgons, "\n") {
		inputChan <- line
	}
	close(inputChan)
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Processing took: %s\n", elapsed)
}
