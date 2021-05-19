package main

import (
	"math/rand"

	"fyne.io/fyne/v2"
	"github.com/AnkushJadhav/fynesweeper/components"
	"github.com/asaskevich/EventBus"
)

type game struct {
	board *fyne.Container
	tiles [][]*components.Tile
}

func newGame(bus EventBus.Bus, rows, cols, mineCount int) *game {
	plan := make([][]int, 0)

	// generate a blank plan
	for itrRow := 0; itrRow < rows; itrRow++ {
		row := make([]int, 0)
		for itrCol := 0; itrCol < cols; itrCol++ {
			row = append(row, 0)
		}
		plan = append(plan, row)
	}

	// generate mines with surrounding info
	plan = generatePlan(plan, mineCount)

	// generate game tiles based on plan
	tiles := make([][]*components.Tile, 0)
	for itrRow := 0; itrRow < rows; itrRow++ {
		row := make([]*components.Tile, 0)
		for itrCol := 0; itrCol < cols; itrCol++ {
			var tile *components.Tile
			switch plan[itrRow][itrCol] {
			case -1:
				tile = components.NewTile(bus, components.TileTypeMine, itrRow, itrCol)
				break
			case 0:
				tile = components.NewTile(bus, components.TileType0, itrRow, itrCol)
				break
			case 1:
				tile = components.NewTile(bus, components.TileType1, itrRow, itrCol)
				break
			case 2:
				tile = components.NewTile(bus, components.TileType2, itrRow, itrCol)
				break
			case 3:
				tile = components.NewTile(bus, components.TileType3, itrRow, itrCol)
				break
			case 4:
				tile = components.NewTile(bus, components.TileType4, itrRow, itrCol)
				break
			case 5:
				tile = components.NewTile(bus, components.TileType5, itrRow, itrCol)
				break
			case 6:
				tile = components.NewTile(bus, components.TileType6, itrRow, itrCol)
				break
			case 7:
				tile = components.NewTile(bus, components.TileType7, itrRow, itrCol)
				break
			case 8:
				tile = components.NewTile(bus, components.TileType8, itrRow, itrCol)
				break
			default:
			}

			row = append(row, tile)
		}
		tiles = append(tiles, row)
	}

	board := components.NewBoard(tiles)

	g := &game{board: board, tiles: tiles}
	return g
}

func generatePlan(plan [][]int, mineCount int) [][]int {
	maxRow := len(plan) - 1
	maxCol := len(plan[0]) - 1
	itr := 0
	for {
		if itr == mineCount {
			break
		}

		r := rand.Intn(maxRow + 1)
		c := rand.Intn(maxCol + 1)
		if isMine(plan, r, c) {
			continue
		}
		plan[r][c] = -1
		if isInBounds(r-1, c-1, 0, 0, maxRow, maxCol) && !isMine(plan, r-1, c-1) {
			plan[r-1][c-1]++
		}
		if isInBounds(r-1, c, 0, 0, maxRow, maxCol) && !isMine(plan, r-1, c) {
			plan[r-1][c]++
		}
		if isInBounds(r-1, c+1, 0, 0, maxRow, maxCol) && !isMine(plan, r-1, c+1) {
			plan[r-1][c+1]++
		}
		if isInBounds(r, c-1, 0, 0, maxRow, maxCol) && !isMine(plan, r, c-1) {
			plan[r][c-1]++
		}
		if isInBounds(r, c+1, 0, 0, maxRow, maxCol) && !isMine(plan, r, c+1) {
			plan[r][c+1]++
		}
		if isInBounds(r+1, c-1, 0, 0, maxRow, maxCol) && !isMine(plan, r+1, c-1) {
			plan[r+1][c-1]++
		}
		if isInBounds(r+1, c, 0, 0, maxRow, maxCol) && !isMine(plan, r+1, c) {
			plan[r+1][c]++
		}
		if isInBounds(r+1, c+1, 0, 0, maxRow, maxCol) && !isMine(plan, r+1, c+1) {
			plan[r+1][c+1]++
		}

		itr++
	}

	return plan
}

func isInBounds(row, col, minRow, minCol, maxRow, maxCol int) bool {
	return row >= minRow && row <= maxRow && col >= minCol && col <= maxCol
}

func isMine(plan [][]int, row, col int) bool {
	return plan[row][col] == -1
}
