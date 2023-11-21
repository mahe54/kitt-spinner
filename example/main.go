package main

import (
	"time"

	"github.com/mahe54/kitt-spinner"
)

func main() {
	//Visualize loading
	done := make(chan bool)
	frames := kitt.PrepareFrames()
	go kitt.Animate(frames, 100*time.Millisecond, done)

	//Simulate work
	time.Sleep(5 * time.Second)

	done <- true
}
