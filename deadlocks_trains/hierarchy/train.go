package hierarchy

import (
	"github.com/lozhkindm/go-multithreading/deadlocks_trains/common"
	"sort"
	"time"
)

func lockIntersections(id, start, end int, cr []*common.Crossing) {
	var toLock []*common.Intersection
	for _, c := range cr {
		if start <= c.Position && c.Position <= end && c.Intersection.LockedBy != id {
			toLock = append(toLock, c.Intersection)
		}
	}

	sort.Slice(toLock, func(i, j int) bool {
		return toLock[i].Id < toLock[j].Id
	})

	for _, is := range toLock {
		is.Mx.Lock()
		is.LockedBy = id
		time.Sleep(10 * time.Millisecond)
	}
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
				c.Intersection.LockedBy = -1
				c.Intersection.Mx.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
