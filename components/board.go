package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// NewBoard creates a new board of dimensions rows and cols
func NewBoard(tiles [][]*Tile) *fyne.Container {
	return container.NewGridWithRows(len(tiles), flattenTiles(tiles)...)
}

func flattenTiles(tiles [][]*Tile) []fyne.CanvasObject {
	res := make([]fyne.CanvasObject, 0)

	for _, row := range tiles {
		for _, item := range row {
			res = append(res, item)
		}
	}

	return res
}
