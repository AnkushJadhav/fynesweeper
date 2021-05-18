package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/AnkushJadhav/fynesweeper/components"
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
	a := app.NewWithID("com.ankush.fynesweeper")
	w := a.NewWindow("fynesweeper")
	a.Settings().SetTheme(&myTheme{})

	ts := make([]fyne.CanvasObject, 0)
	for itr := 0; itr < 100; itr++ {
		ts = append(ts, components.NewTile(components.TileTypeMineActive))
	}

	c := container.NewGridWithColumns(10, ts...)

	w.SetContent(c)
	w.ShowAndRun()
}
