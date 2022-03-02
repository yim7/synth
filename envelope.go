package main

type Phase int

const (
	Inactive Phase = iota
	Attack
	Decay
	Sustain
	Release
)

type Envelope struct {
	sampleRate      float64
	attack          float64
	decay           float64
	sustain         float64
	sustainLevel    float64
	release         float64
	phase           Phase
	hold            bool
	sampleProcessed int

	input Processor
}

func NewEnvelope(attack, decay, sustainLevel, release, sampleRate float64) *Envelope {

	return &Envelope{
		attack:       attack,
		decay:        decay,
		sustainLevel: sustainLevel,
		release:      release,
		sampleRate:   sampleRate,
		phase:        Inactive,
		hold:         true,
	}
}

func (e *Envelope) Channels() int {
	return 1
}

func (e *Envelope) NoteOn() {
	e.sampleProcessed = 0
	e.hold = true
	e.phase = Attack
}

func (e *Envelope) NoteOff() {
	e.hold = false
}

func (e *Envelope) SetInput(input Processor) {
	e.input = input
}

func (e *Envelope) Process(frames int) []float64 {
	samples := e.input.Process(frames)
	for i, v := range samples {
		samples[i] = v * e.Value()
	}

	return samples
}

func (e *Envelope) Value() (amplitude float64) {
	sampleRate := e.sampleRate
	attackSamples := e.attack * sampleRate
	decaySamples := e.decay * sampleRate
	releaseSamples := e.release * sampleRate

	switch e.phase {
	case Inactive:
	case Attack:
		amplitude = float64(e.sampleProcessed-0) / (attackSamples - 0)
		if amplitude >= 1 {
			e.phase = Decay
			e.sampleProcessed = 0
		}
	case Decay:
		amplitude = 1 + (e.sustainLevel-1)*float64(e.sampleProcessed)/decaySamples
		if amplitude <= e.sustainLevel {
			e.phase = Sustain
		}
	case Sustain:
		amplitude = e.sustainLevel
		if !e.hold {
			e.phase = Release
			e.sampleProcessed = 0
		}
	case Release:
		if e.sampleProcessed <= int(releaseSamples) {
			amplitude = e.sustainLevel + (0-e.sustainLevel)*float64(e.sampleProcessed)/releaseSamples
		} else {
			e.phase = Inactive
		}
	}
	e.sampleProcessed++
	// fmt.Println("amplitude:", amplitude)
	return
}
