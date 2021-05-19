package main

import (
	"fyne.io/fyne/v2"
	"github.com/AnkushJadhav/fynesweeper/components"
	"github.com/AnkushJadhav/fynesweeper/events"
	"github.com/asaskevich/EventBus"
)

func initEngine(game *game, bus EventBus.Bus, w fyne.Window) {
	bus.Subscribe(events.EventTileOpened, tileOpenHandler(game.tiles))
	bus.Subscribe(events.EventSmileyManTriggered, smileyMantriggerHandler(bus, w))
}

func smileyMantriggerHandler(bus EventBus.Bus, w fyne.Window) func() {
	return func() {
		game := newGame(bus, 20, 20, 20)
		w.SetContent(game.board)
	}
}

func tileOpenHandler(tiles [][]*components.Tile) func(int, int) {
	return func(row, col int) {
		switch tiles[row][col].Base {
		case components.TileTypeMine:
			tiles[row][col].Open(true)
			endGame(tiles)
			break
		case components.TileType0:
			revealEdges(tiles, row, col)
			break
		default:
			tiles[row][col].Open(true)
		}
		return
	}
}

func endGame(tiles [][]*components.Tile) {
	for itrRow := 0; itrRow < len(tiles); itrRow++ {
		for itrCol := 0; itrCol < len(tiles[itrRow]); itrCol++ {
			if !tiles[itrRow][itrCol].IsOpen {
				tiles[itrRow][itrCol].Open(false)
			}
		}
	}
}

// 8-directional flood fill
func revealEdges(tiles [][]*components.Tile, row, col int) {
	if row < 0 || col < 0 || row >= len(tiles) || col >= len(tiles[0]) {
		return
	}
	if tiles[row][col].IsOpen {
		return
	}
	if tiles[row][col].Base != components.TileType0 && tiles[row][col].Base != components.TileTypeMine {
		tiles[row][col].Open(false)
		return
	}
	tiles[row][col].Open(false)
	revealEdges(tiles, row-1, col-1)
	revealEdges(tiles, row-1, col)
	revealEdges(tiles, row-1, col+1)
	revealEdges(tiles, row, col-1)
	revealEdges(tiles, row, col+1)
	revealEdges(tiles, row+1, col-1)
	revealEdges(tiles, row+1, col)
	revealEdges(tiles, row+1, col+1)
	return
}
