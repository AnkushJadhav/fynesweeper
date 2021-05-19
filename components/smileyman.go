package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/AnkushJadhav/fynesweeper/events"
	"github.com/asaskevich/EventBus"
)

// The different states of a game
const (
	// Are ya winnin' son?
	GameStateWin = iota
	GameStateLose
	GameStateOngoing
)

// GameState represents the current state of the game
type GameState int

// SmileyMan controls everything. God, you mean? No.
type SmileyMan struct {
	widget.Icon

	gameState GameState

	bus EventBus.Bus
}

// NewSmileyMan creates a new smileyman
func NewSmileyMan(bus EventBus.Bus, state GameState) *SmileyMan {
	sm := &SmileyMan{bus: bus, gameState: state}
	sm.ExtendBaseWidget(sm)
	sm.SetResource(getResourceByGameState(state))

	return sm
}

// SetState chagnes smiley amns state based on game state
func (sm *SmileyMan) SetState(state GameState) {
	sm.SetResource(getResourceByGameState(state))
}

// MouseDown impl for desktop clicks
func (sm *SmileyMan) MouseDown(ev *desktop.MouseEvent) {
	if ev.Button == desktop.MouseButtonPrimary {
		sm.SetResource(resourceFacepressedPng)
	}
}

// MouseUp impl for desktop clicks
func (sm *SmileyMan) MouseUp(ev *desktop.MouseEvent) {
	if ev.Button == desktop.MouseButtonPrimary {
		sm.SetResource(getResourceByGameState(GameStateOngoing))
		sm.bus.Publish(events.EventSmileyManTriggered)
	}
}

func getResourceByGameState(state GameState) fyne.Resource {
	var res fyne.Resource

	switch state {
	case GameStateLose:
		res = resourceFacelosePng
		break
	case GameStateWin:
		res = resourceFacewinPng
		break
	case GameStateOngoing:
		res = resourceFaceunpressedPng
		break
	default:
	}

	return res
}
