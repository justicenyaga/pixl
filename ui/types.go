package ui

import (
	"fyne.io/fyne/v2"
	"justicenyaga.io/pixl/apptype"
	"justicenyaga.io/pixl/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
