package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// TileTypes
const (
	TileTypeMine = iota
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

	Base      TileType
	IsOpen    bool
	IsFlagged bool
	Row       int
	Col       int

	OpenHandler func(int, int)
}

// NewTile creates a new tile of the given type
func NewTile(tileType TileType, row, col int, openHandler func(int, int)) *Tile {
	t := &Tile{Base: tileType, IsOpen: false, Row: row, Col: col, OpenHandler: openHandler}
	t.ExtendBaseWidget(t)
	t.SetResource(resourceClosedPng)
	return t
}

// Open opens a tile
func (t *Tile) Open(byUser bool) {
	if !t.IsOpen {
		t.open(byUser)
	}
}

// MouseIn impl for desktop clicks
func (t *Tile) MouseIn(ev *desktop.MouseEvent) {
	if ev.Button == desktop.MouseButtonPrimary && !t.IsOpen && !t.IsFlagged {
		t.SetResource(resourceType0Png)
	}
}

// MouseOut impl for desktop clicks
func (t *Tile) MouseOut() {
	if !t.IsOpen && !t.IsFlagged {
		t.SetResource(resourceClosedPng)
	}
}

// MouseMoved impl for desktop clicks
func (t *Tile) MouseMoved(ev *desktop.MouseEvent) {
	// do nothing
}

// MouseDown impl for desktop clicks
func (t *Tile) MouseDown(ev *desktop.MouseEvent) {
	if ev.Button == desktop.MouseButtonPrimary && !t.IsOpen && !t.IsFlagged {
		t.SetResource(resourceType0Png)
	}
}

// MouseUp impl for desktop clicks
func (t *Tile) MouseUp(ev *desktop.MouseEvent) {
	if ev.Button == desktop.MouseButtonPrimary && !t.IsOpen && !t.IsFlagged {
		t.OpenHandler(t.Row, t.Col)
	}
	if ev.Button == desktop.MouseButtonSecondary {
		if !t.IsFlagged {
			t.flag()
		}
	}
}

func (t *Tile) flag() {
	t.SetResource(resourceFlagPng)
	t.IsFlagged = true
}

func (t *Tile) open(byUser bool) {
	t.SetResource(getBaseResourceByType(t.Base, byUser))
	t.IsOpen = true
}

func getBaseResourceByType(t TileType, byUser bool) fyne.Resource {
	var res fyne.Resource

	switch t {
	case TileTypeMine:
		if byUser {
			res = resourceMineredPng
		} else {
			res = resourceMinePng
		}
	case TileType0:
		res = resourceType0Png
	case TileType1:
		res = resourceType1Png
	case TileType2:
		res = resourceType2Png
	case TileType3:
		res = resourceType3Png
	case TileType4:
		res = resourceType4Png
	case TileType5:
		res = resourceType5Png
	case TileType6:
		res = resourceType6Png
	case TileType7:
		res = resourceType7Png
	case TileType8:
		res = resourceType8Png
	}
	return res
}
