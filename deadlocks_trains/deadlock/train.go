package deadlock

import (
	"github.com/lozhkindm/go-multithreading/deadlocks_trains/common"
	"time"
)

func MoveTrain(t *common.Train, dist int, cr []*common.Crossing) {
	for t.Front < dist {
		t.Front += 1
		for _, c := range cr {
			if t.Front == c.Position {
				c.Intersection.Mx.Lock()
				c.Intersection.LockedBy = t.Id
			}
			back := t.Front - t.Length
			if back == c.Position {
				c.Intersection.LockedBy = -1
				c.Intersection.Mx.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
