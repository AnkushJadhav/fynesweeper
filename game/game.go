package game

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/AnkushJadhav/fynesweeper/components"
)

// Game is a game damnit
type Game struct {
	Board       *fyne.Container
	Tiles       [][]*components.Tile
	Smiley      *components.SmileyMan
	MineCounter *components.MineCounter

	OpenCount int
	WinCount  int

	Win fyne.Window
}

// NewGame creates a new game
func NewGame() *Game {
	return &Game{}
}

// SeedGame creates a new game damnit
func (g *Game) SeedGame(rows, cols, mineCount int) {
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
				tile = components.NewTile(components.TileTypeMine, itrRow, itrCol, g.openTile)
				break
			case 0:
				tile = components.NewTile(components.TileType0, itrRow, itrCol, g.openTile)
				break
			case 1:
				tile = components.NewTile(components.TileType1, itrRow, itrCol, g.openTile)
				break
			case 2:
				tile = components.NewTile(components.TileType2, itrRow, itrCol, g.openTile)
				break
			case 3:
				tile = components.NewTile(components.TileType3, itrRow, itrCol, g.openTile)
				break
			case 4:
				tile = components.NewTile(components.TileType4, itrRow, itrCol, g.openTile)
				break
			case 5:
				tile = components.NewTile(components.TileType5, itrRow, itrCol, g.openTile)
				break
			case 6:
				tile = components.NewTile(components.TileType6, itrRow, itrCol, g.openTile)
				break
			case 7:
				tile = components.NewTile(components.TileType7, itrRow, itrCol, g.openTile)
				break
			case 8:
				tile = components.NewTile(components.TileType8, itrRow, itrCol, g.openTile)
				break
			default:
			}

			row = append(row, tile)
		}
		tiles = append(tiles, row)
	}

	board := components.NewBoard(tiles)
	sm := components.NewSmileyMan(components.GameStateOngoing, g.resetHandler)

	mc := components.NewMineCounter(mineCount)

	g.Board = board
	g.Tiles = tiles
	g.Smiley = sm
	g.MineCounter = mc
	g.OpenCount = 0
	g.WinCount = (rows * cols) - mineCount
}

func generatePlan(plan [][]int, mineCount int) [][]int {
	rand.Seed(time.Now().UnixNano())
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

// Render the game
func (g *Game) Render() {
	t := container.NewVBox(container.NewHBox(g.MineCounter.Container, g.Smiley), g.Board)

	g.Win.SetContent(t)
}

func (g *Game) resetHandler() {
	g.SeedGame(20, 20, 1)
	g.Render()
}

func isInBounds(row, col, minRow, minCol, maxRow, maxCol int) bool {
	return row >= minRow && row <= maxRow && col >= minCol && col <= maxCol
}

func isMine(plan [][]int, row, col int) bool {
	return plan[row][col] == -1
}
