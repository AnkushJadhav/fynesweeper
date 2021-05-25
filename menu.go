package main

import (
	"fyne.io/fyne/v2"
	"github.com/AnkushJadhav/fynesweeper/game"
)

func newMenuBar(g *game.Game) *fyne.MainMenu {
	miSizeBeginner := fyne.NewMenuItem("Beginner", func() {
		g.SeedGame(game.SizeBeginner)
		g.Render()
	})
	miSizeIntermediate := fyne.NewMenuItem("Intermediate", func() {
		g.SeedGame(game.SizeIntermediate)
		g.Render()
	})
	miSizeExpert := fyne.NewMenuItem("Expert", func() {
		g.SeedGame(game.SizeExpert)
		g.Render()
	})

	mDiffculty := fyne.NewMenu("Difficulty", miSizeBeginner, miSizeIntermediate, miSizeExpert)
	mm := fyne.NewMainMenu(mDiffculty)
	return mm
}
