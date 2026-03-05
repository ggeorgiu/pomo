package main

import (
	"sync"
	"time"
)

type progressBar struct {
	milliseconds int
	mutex        sync.Mutex
	percentage   float32
	tick         chan struct{}
}

func newProgressBar(seconds int) (*progressBar, chan struct{}) {
	tick := make(chan struct{})

	return &progressBar{
		milliseconds: seconds * 1000,
		tick:         tick,
	}, tick
}

func (pb *progressBar) run() {
	ticker := time.Tick(100 * time.Millisecond)
	deadline := time.Now().Add(time.Duration(pb.milliseconds) * time.Millisecond)

	ratio := float32(100) / float32(pb.milliseconds) * 100

	for time.Now().Before(deadline) {
		<-ticker

		pb.part(ratio)
		pb.tick <- struct{}{}
	}

	close(pb.tick)
}

func (pb *progressBar) part(p float32) {
	pb.mutex.Lock()
	defer pb.mutex.Unlock()
	pb.percentage += p
}

func (pb *progressBar) quant() float32 {
	pb.mutex.Lock()
	defer pb.mutex.Unlock()

	return pb.percentage
}
