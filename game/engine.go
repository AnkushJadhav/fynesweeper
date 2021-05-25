package game

import (
	"github.com/AnkushJadhav/fynesweeper/components"
)

func (g *Game) openTile(row, col int) {
	switch g.Tiles[row][col].Base {
	case components.TileTypeMine:
		g.Tiles[row][col].Open(true)
		g.lose()
		break
	case components.TileType0:
		g.revealEdges(row, col)
		break
	default:
		g.Tiles[row][col].Open(true)
		g.OpenCount++
	}
	if g.OpenCount == g.WinCount {
		g.win()
	}
	return
}

func (g *Game) win() {
	for itrRow := 0; itrRow < len(g.Tiles); itrRow++ {
		for itrCol := 0; itrCol < len(g.Tiles[itrRow]); itrCol++ {
			if !g.Tiles[itrRow][itrCol].IsOpen {
				g.Tiles[itrRow][itrCol].Flag()
			}
		}
	}
	g.Smiley.SetState(components.GameStateWin)
}

func (g *Game) lose() {
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
func (g *Game) revealEdges(row, col int) {
	if row < 0 || col < 0 || row >= len(g.Tiles) || col >= len(g.Tiles[0]) {
		return
	}
	if g.Tiles[row][col].IsOpen {
		return
	}

	if g.Tiles[row][col].Base != components.TileType0 && g.Tiles[row][col].Base != components.TileTypeMine {
		g.Tiles[row][col].Open(false)
		g.OpenCount++
		return
	}
	g.Tiles[row][col].Open(false)
	g.OpenCount++
	g.revealEdges(row-1, col-1)
	g.revealEdges(row-1, col)
	g.revealEdges(row-1, col+1)
	g.revealEdges(row, col-1)
	g.revealEdges(row, col+1)
	g.revealEdges(row+1, col-1)
	g.revealEdges(row+1, col)
	g.revealEdges(row+1, col+1)
	return
}
