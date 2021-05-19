package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/AnkushJadhav/fynesweeper/components"
	"github.com/asaskevich/EventBus"
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
	bus := EventBus.New()
	w := a.NewWindow("fynesweeper")

	game := newGame(bus, 20, 20, 20)
	initEngine(game, bus, w)
	sm := components.NewSmileyMan(bus, components.GameStateOngoing)

	t := container.NewVBox(sm, game.board)

	w.SetContent(t)
	w.ShowAndRun()
}
