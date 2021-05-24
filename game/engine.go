package game

import (
	"github.com/AnkushJadhav/fynesweeper/components"
)

func (g *Game) openTile(row, col int) {
	switch g.Tiles[row][col].Base {
	case components.TileTypeMine:
		g.Tiles[row][col].Open(true)
		g.end()
		break
	case components.TileType0:
		revealEdges(g.Tiles, row, col)
		break
	default:
		g.Tiles[row][col].Open(true)
	}
	return
}

func (g *Game) end() {
	for itrRow := 0; itrRow < len(g.Tiles); itrRow++ {
		for itrCol := 0; itrCol < len(g.Tiles[itrRow]); itrCol++ {
			if !g.Tiles[itrRow][itrCol].IsOpen {
				g.Tiles[itrRow][itrCol].Open(false)
			}
		}
	}
	g.Smiley.SetState(components.GameStateLose)
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
