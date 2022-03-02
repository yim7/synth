package main

type Harmonics struct {
	weight []float64
	oscis  []*Oscillator
}

func NewHarmonics(frequency, sampleRate float64, weight []float64) *Harmonics {
	size := len(weight)

	os := make([]*Oscillator, size)

	for i := 0; i < size; i++ {
		o := NewOscillator(frequency*float64(i+1), sampleRate, Sine)
		os[i] = o
	}

	return &Harmonics{
		weight: weight,
		oscis:  os,
	}
}

func (e *Harmonics) SetFrequency(f float64) {
	for i, v := range e.oscis {
		v.frequency = f * float64(i+1)
	}
}

func (e *Harmonics) Channels() int {
	return 1
}

func (e *Harmonics) Process(frames int) []float64 {
	size := len(e.weight)
	var sum float64
	for _, w := range e.weight {
		sum += w
	}

	samplesList := make([][]float64, size)

	result := make([]float64, frames)

	for i := 0; i < size; i++ {
		samplesList[i] = e.oscis[i].Process(frames)
	}

	for n := 0; n < frames; n++ {
		var sample float64
		for osciIndex := 0; osciIndex < size; osciIndex++ {
			sample += samplesList[osciIndex][n] * e.weight[osciIndex] / sum
		}
		result[n] = sample
	}

	return result
}
