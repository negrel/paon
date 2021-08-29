package metrics

import "time"

type durationSlice struct {
	min, max time.Duration
	data     []time.Duration
}

func newDurationSlice() *durationSlice {
	return &durationSlice{
		min:  time.Duration(time.Hour),
		data: make([]time.Duration, 0, 65535),
	}
}

func (ds *durationSlice) push(duration time.Duration) {
	ds.data = append(ds.data, duration)

	if duration < ds.min {
		ds.min = duration
	} else if duration > ds.max {
		ds.max = duration
	}
}

func (ds *durationSlice) mean() time.Duration {
	var result int
	for _, duration := range ds.data {
		result += int(duration)
	}

	return time.Duration(result / len(ds.data))
}
