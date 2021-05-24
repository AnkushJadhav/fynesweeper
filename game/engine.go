package game

import (
	"fyne.io/fyne/v2/container"
	"github.com/AnkushJadhav/fynesweeper/components"
)

// Render the game
func (g *Game) Render() {
	t := container.NewVBox(g.Smiley, g.Board)

	g.Win.SetContent(t)
}

func (g *Game) resetHandler() {
	g.SeedGame(20, 20, 20)
	g.Render()
}

func (g *Game) openTile(row, col int) {
	switch g.Tiles[row][col].Base {
	case components.TileTypeMine:
		g.Tiles[row][col].Open(true)
		endGame(g.Tiles)
		break
	case components.TileType0:
		revealEdges(g.Tiles, row, col)
		break
	default:
		g.Tiles[row][col].Open(true)
	}
	return
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
