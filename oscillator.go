package main

import (
	"math"
	"time"
)

// 0 <= step <= 1
type WaveFormGenerator func(step float64) float64

func Sine(step float64) float64 {
	return math.Sin(2 * math.Pi * step)
}

func Square(step float64) float64 {
	if step < 0.5 {
		return -1
	} else {
		return 1
	}
}

func Sawtooth(step float64) float64 {
	return 2 * (step - math.Round(step))
}

func Triangle(step float64) float64 {
	return 1 - 4*math.Abs(step-math.Round(step))
}

type Oscillator struct {
	sampleRate float64
	frequency  float64
	waveForm   WaveFormGenerator
	waveTable  []float64
	phase      float64
}

func NewOscillator(frequency, sampleRate float64, waveForm WaveFormGenerator) *Oscillator {
	o := &Oscillator{
		frequency:  frequency,
		sampleRate: sampleRate,
		waveForm:   waveForm,
	}

	o.fillWaveTable()

	return o
}

func (o *Oscillator) fillWaveTable() {
	f := o.waveForm
	sampleRate := o.sampleRate
	waveSize := int(sampleRate)
	waveTable := make([]float64, waveSize)
	o.waveTable = waveTable

	for i := 0; i < waveSize; i++ {
		waveTable[i] = f(float64(i) / sampleRate)
	}
}

func (o *Oscillator) Process(frames int) []float64 {
	waveTable := o.waveTable
	waveSize := len(waveTable)

	freq := o.frequency
	sampleRate := o.sampleRate
	incr := float64(waveSize) / sampleRate * freq
	phase := o.phase

	samples := make([]float64, frames)
	for i := 0; i < frames; i++ {
		index := int(phase)
		samples[i] = waveTable[index]
		phase += incr
		phase -= math.Floor(phase/float64(waveSize)) * float64(waveSize)
	}
	o.phase = phase

	return samples
}

func (o *Oscillator) Channels() int {
	return 1
}

func (o *Oscillator) ProcessDuraion(duration time.Duration) []float64 {
	frames := int(duration.Seconds() * o.sampleRate)
	return o.Process(frames)
}
