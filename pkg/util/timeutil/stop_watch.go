package timeutil

import (
	"time"
)

type StopWatch struct {
	start time.Time
	end   time.Time
}

func (stopWatch *StopWatch) Start() {
	stopWatch.start = time.Now()
}

func (stopWatch *StopWatch) Stop() {
	stopWatch.end = time.Now()
}

func (stopWatch *StopWatch) GetDurationInMilliseconds() int64 {
	if stopWatch.start.After(stopWatch.end) {
		return -1
	}

	return stopWatch.end.Sub(stopWatch.start).Milliseconds()
}
