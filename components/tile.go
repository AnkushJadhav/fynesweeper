package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// TileTypes
const (
	TileTypeMineActive = iota
	TileTypeMineInactive
	TileTypeEmpty
	TileType0
	TileType1
	TileType2
	TileType3
	TileType4
	TileType5
	TileType6
	TileType7
	TileType8
)

// TileType is the type of the tile that modifies it's look
type TileType int

// Tile stores the data and state of a tile on the game board
type Tile struct {
	widget.Icon

	base   TileType
	isOpen bool
}

// NewTile creates a new tile of the given type
func NewTile(tileType TileType) *Tile {
	t := &Tile{base: tileType, isOpen: false}
	t.ExtendBaseWidget(t)
	t.SetResource(resourceClosedPng)
	return t
}

// MouseIn impl for desktop clicks
func (t *Tile) MouseIn(ev *desktop.MouseEvent) {
	if ev.Button == desktop.MouseButtonPrimary && !t.isOpen {
		t.SetResource(resourceType0Png)
	}
}

// MouseOut impl for desktop clicks
func (t *Tile) MouseOut() {
	if !t.isOpen {
		t.SetResource(resourceClosedPng)
	}
}

// MouseMoved impl for desktop clicks
func (t *Tile) MouseMoved(ev *desktop.MouseEvent) {
	// do nothing
}

// MouseDown impl for desktop clicks
func (t *Tile) MouseDown(ev *desktop.MouseEvent) {
	if !t.isOpen {
		t.SetResource(resourceType0Png)
	}
}

// MouseUp impl for desktop clicks
func (t *Tile) MouseUp(ev *desktop.MouseEvent) {
	if !t.isOpen {
		t.SetResource(getBaseResourceByType(t.base))
		t.isOpen = true
	}
}

func getBaseResourceByType(t TileType) fyne.Resource {
	var res fyne.Resource

	switch t {
	case TileTypeMineActive:
		res = resourceMineredPng
	}
	return res
}
