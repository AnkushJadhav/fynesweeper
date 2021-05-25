package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/AnkushJadhav/fynesweeper/game"
)

const (
	appID      = "me.ankushjadhav.fynesweeper"
	appVersion = "0.0.1"
	appName    = "fynesweeper"
)

func main() {
	// create a fyne app
	a := app.NewWithID(appID)
	a.SetIcon(resourceIconPng)
	a.Settings().SetTheme(newCustomTheme())
	w := a.NewWindow(appName)

	// create a new default game
	g := game.NewGame()
	g.SeedGame(game.SizeBeginner)

	// render the game
	w.SetMainMenu(newMenuBar(g))
	g.Win = w
	g.Render()
	w.ShowAndRun()
}
