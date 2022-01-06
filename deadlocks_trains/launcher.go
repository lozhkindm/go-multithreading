package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/lozhkindm/go-multithreading/deadlocks_trains/arbitrator"
	"github.com/lozhkindm/go-multithreading/deadlocks_trains/common"
	"log"
	"sync"
)

var (
	trains [4]*common.Train
	isecs  [4]*common.Intersection
)

const (
	TrainLength = 70
)

func update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		DrawTracks(screen)
		DrawIntersections(screen)
		DrawTrains(screen)
	}
	return nil
}

func main() {
	for i := 0; i < 4; i++ {
		trains[i] = &common.Train{Id: i, Length: TrainLength, Front: 0}
		isecs[i] = &common.Intersection{Id: i, Mx: sync.Mutex{}, LockedBy: -1}
	}

	go arbitrator.MoveTrain(
		trains[0],
		300,
		[]*common.Crossing{
			{Position: 125, Intersection: isecs[0]},
			{Position: 175, Intersection: isecs[1]},
		},
	)

	go arbitrator.MoveTrain(
		trains[1],
		300,
		[]*common.Crossing{
			{Position: 125, Intersection: isecs[1]},
			{Position: 175, Intersection: isecs[2]},
		},
	)

	go arbitrator.MoveTrain(
		trains[2],
		300,
		[]*common.Crossing{
			{Position: 125, Intersection: isecs[2]},
			{Position: 175, Intersection: isecs[3]},
		},
	)

	go arbitrator.MoveTrain(
		trains[3],
		300,
		[]*common.Crossing{
			{Position: 125, Intersection: isecs[3]},
			{Position: 175, Intersection: isecs[0]},
		},
	)

	if err := ebiten.Run(update, 320, 320, 3, "Trains in a box"); err != nil {
		log.Fatal(err)
	}
}
