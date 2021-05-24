package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/AnkushJadhav/fynesweeper/game"
)

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Font(name fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(name)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNamePadding {
		return 0
	}
	return theme.DefaultTheme().Size(name)
}

func main() {
	a := app.NewWithID("me.ankushjadhav.fynesweeper")
	a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("fynesweeper")

	g := game.NewGame()
	g.SeedGame(20, 20, 20)
	g.Win = w
	g.Render()

	w.ShowAndRun()
}
