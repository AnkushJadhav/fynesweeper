package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/AnkushJadhav/fynesweeper/game"
)

const (
	appID          = "me.ankushjadhav.fynesweeper"
	appVersion     = "0.0.1"
	appWindowTitle = "fynesweeper"
)

func main() {
	// create a fyne app
	a := app.NewWithID(appID)
	a.Settings().SetTheme(newCustomTheme())
	w := a.NewWindow(appWindowTitle)

	// create a new default game
	g := game.NewGame()
	g.SeedGame(game.SizeBeginner)

	// render the game
	g.Win = w
	g.Render()
	w.ShowAndRun()
}
