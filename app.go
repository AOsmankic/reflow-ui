package main

import (
	"context"
	"math/rand"
)

type Coordinate struct {
	X float32
	Y float32
}

type State struct {
	Temp            float32
	SelectedProfile string
	SecondsLeft     int
}

var tempBuffer []Coordinate

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetState() State {
	return State{
		Temp:            rand.Float32(),
		SelectedProfile: SelectedProfile.Name,
		SecondsLeft:     156,
	}
}
func (a *App) Init() {

	LoadProfiles()
}

func (a *App) GetSelectedProfile() Profile {
	return SelectedProfile
}

func (a *App) GetAllProfiles() []Profile {
	return GetProfilesList()
}

func (a *App) GetTempHistory() []Coordinate {
	// TODO: Get Temp Data From Doug
	var xCoord float32
	if len(tempBuffer) > 0 {
		xCoord = tempBuffer[len(tempBuffer)-1].X + 1
	} else {
		xCoord = 1
	}

	tempBuffer = append(tempBuffer, Coordinate{X: xCoord, Y: rand.Float32()})
	return tempBuffer
}

func (a *App) SetProfile(profileID string) {
	profile := GetProfile(profileID)
	setProfile(profile)
}

func (a *App) ResetBuffer() bool {
	tempBuffer = []Coordinate{}
	return true
}

func (a *App) ShiftTempBuffer(scale float32) {
	var finalBuffer []Coordinate
	for _, coordinate := range tempBuffer {
		coordinate.X = coordinate.X + scale
		finalBuffer = append(finalBuffer, coordinate)
	}
	tempBuffer = finalBuffer
}

func (a *App) GetTempCurve(profileID string) []Coordinate {
	return GetProfile(profileID).CurveValues
}

func (a *App) FindMaxY() float32 {
	var highestY float32 = -1.0
	for _, coordinate := range tempBuffer {
		if coordinate.Y > highestY {
			highestY = coordinate.Y
		}
	}
	return highestY
}

func (a *App) FindMaxX() float32 {
	var highestX float32 = -1.0
	for _, coordinate := range tempBuffer {
		if coordinate.X > highestX {
			highestX = coordinate.X
		}
	}
	return highestX
}
