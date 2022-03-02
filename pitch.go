package main

type Pitch float64

var PitchTable = map[string]Pitch{
	"A3": 220,
	"C4": 261,
	"D4": 293,
	"E4": 329,
	"F4": 349,
	"G4": 392,
	"A4": 440,
	"B4": 493,
}
