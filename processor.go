package main

type Processor interface {
	Process(frames int) []float64
	// ProcessDuraion(time.Duration) []float64
	// Channels() int
}
