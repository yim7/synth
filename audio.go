package main

import (
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type AudioPlayer struct {
	device     sdl.AudioDeviceID
	SampleRate int
	Channels   int
	input      Processor
}

func OpenAudioPlayer(sampleRate int, channels int) (*AudioPlayer, error) {
	spec := &sdl.AudioSpec{
		Freq:     int32(sampleRate),
		Format:   sdl.AUDIO_U8,
		Channels: uint8(channels),
		Samples:  4096,
	}

	device, err := sdl.OpenAudioDevice("", false, spec, nil, sdl.AUDIO_ALLOW_ANY_CHANGE)
	if err != nil {
		return nil, err
	}

	return &AudioPlayer{
		device:     device,
		SampleRate: sampleRate,
		Channels:   channels,
	}, nil
}

func (player *AudioPlayer) Resume() {
	sdl.PauseAudioDevice(player.device, false)
}

func (player *AudioPlayer) Pause() {
	sdl.PauseAudioDevice(player.device, true)
}

func (player *AudioPlayer) SetInput(input Processor) {
	player.input = input
}

func (player *AudioPlayer) sendToDevice(samples []float64) {
	data := make([]byte, len(samples))
	for i := range data {
		data[i] = byte((samples[i] + 1) * 255 / 2)
	}

	if err := sdl.QueueAudio(player.device, data); err != nil {
		log.Println(err)
	}
}

func (player *AudioPlayer) Run() {
	duration := 100 * time.Millisecond
	sleep := duration / 100
	sampleFrames := int(duration.Seconds() * float64(player.SampleRate))
	reserveSize := uint32(sampleFrames * 2)
	log.Println("reserve size:", reserveSize)

	for {
		if sdl.GetQueuedAudioSize(player.device) > reserveSize {
			time.Sleep(sleep)
		} else {
			data := player.input.Process(sampleFrames)
			player.sendToDevice(data)
		}
	}
}
