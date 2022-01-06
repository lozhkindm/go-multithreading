package arbitrator

import (
	"github.com/lozhkindm/go-multithreading/deadlocks_trains/common"
	"sync"
	"time"
)

var (
	controller = sync.Mutex{}
	cond       = sync.NewCond(&controller)
)

func allFree(its []*common.Intersection) bool {
	for _, it := range its {
		if it.LockedBy >= 0 {
			return false
		}
	}
	return true
}

func lockIntersections(id, start, end int, cr []*common.Crossing) {
	var toLock []*common.Intersection
	for _, c := range cr {
		if start <= c.Position && c.Position <= end && c.Intersection.LockedBy != id {
			toLock = append(toLock, c.Intersection)
		}
	}

	controller.Lock()
	for !allFree(toLock) {
		cond.Wait()
	}
	for _, is := range toLock {
		is.LockedBy = id
		time.Sleep(10 * time.Millisecond)
	}
	controller.Unlock()
}

func MoveTrain(t *common.Train, dist int, cr []*common.Crossing) {
	for t.Front < dist {
		t.Front += 1
		for _, c := range cr {
			if t.Front == c.Position {
				lockIntersections(t.Id, c.Position, c.Position+t.Length, cr)
			}
			back := t.Front - t.Length
			if back == c.Position {
				controller.Lock()
				c.Intersection.LockedBy = -1
				cond.Broadcast()
				controller.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
