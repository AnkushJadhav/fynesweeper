package main

import (
	"fmt"

	"github.com/AnkushJadhav/fynesweeper/components"
	"github.com/AnkushJadhav/fynesweeper/events"
	"github.com/asaskevich/EventBus"
)

func initEngine(game *game, bus EventBus.Bus) {
	bus.Subscribe(events.EventTileOpened, tileOpenHandler(game.tiles))
}

func tileOpenHandler(game [][]*components.Tile) func(int, int) {
	return func(row, col int) {
		switch game[row][col].Base {
		case components.TileTypeMine:
			fmt.Println("PHATAAA!")
			break
		case components.TileType0:
			fmt.Println("DO unveiling logic!")
			break
		}
		return
	}
}
