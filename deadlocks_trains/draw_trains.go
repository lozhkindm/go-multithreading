package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/lozhkindm/go-multithreading/deadlocks_trains/common"
	"image/color"
	"math"
)

var (
	colors = [4]color.RGBA{
		{R: 233, G: 33, B: 40, A: 255},
		{R: 78, G: 151, B: 210, A: 255},
		{R: 251, G: 170, B: 26, A: 255},
		{R: 11, G: 132, B: 54, A: 255},
	}

	white = color.RGBA{R: 185, G: 185, B: 185, A: 255}
)

func DrawIntersections(screen *ebiten.Image) {
	drawIntersection(screen, isecs[0], 145, 145)
	drawIntersection(screen, isecs[1], 175, 145)
	drawIntersection(screen, isecs[2], 175, 175)
	drawIntersection(screen, isecs[3], 145, 175)
}

func DrawTracks(screen *ebiten.Image) {
	for i := 0; i < 300; i++ {
		screen.Set(10+i, 135, white)
		screen.Set(185, 10+i, white)
		screen.Set(310-i, 185, white)
		screen.Set(135, 310-i, white)
	}
}

func DrawTrains(screen *ebiten.Image) {
	drawXTrain(screen, 0, 1, 10, 135)
	drawYTrain(screen, 1, 1, 10, 185)
	drawXTrain(screen, 2, -1, 310, 185)
	drawYTrain(screen, 3, -1, 310, 135)
}

func drawIntersection(screen *ebiten.Image, intersection *common.Intersection, x int, y int) {
	c := white
	if intersection.LockedBy >= 0 {
		c = colors[intersection.LockedBy]
	}
	screen.Set(x-1, y, c)
	screen.Set(x, y-1, c)
	screen.Set(x, y, c)
	screen.Set(x+1, y, c)
	screen.Set(x, y+1, c)
}

func drawXTrain(screen *ebiten.Image, id int, dir int, start int, yPos int) {
	s := start + (dir * (trains[id].Front - trains[id].Length))
	e := start + (dir * trains[id].Front)
	for i := math.Min(float64(s), float64(e)); i <= math.Max(float64(s), float64(e)); i++ {
		screen.Set(int(i)-dir, yPos-1, colors[id])
		screen.Set(int(i), yPos, colors[id])
		screen.Set(int(i)-dir, yPos+1, colors[id])
	}
}

func drawYTrain(screen *ebiten.Image, id int, dir int, start int, xPos int) {
	s := start + (dir * (trains[id].Front - trains[id].Length))
	e := start + (dir * trains[id].Front)
	for i := math.Min(float64(s), float64(e)); i <= math.Max(float64(s), float64(e)); i++ {
		screen.Set(xPos-1, int(i)-dir, colors[id])
		screen.Set(xPos, int(i), colors[id])
		screen.Set(xPos+1, int(i)-dir, colors[id])
	}
}
