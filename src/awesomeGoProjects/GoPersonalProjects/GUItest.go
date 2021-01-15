package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {

	goApp := app.New()
	lightTheme := theme.LightTheme()
	goApp.Settings().SetTheme(lightTheme)

	w := goApp.NewWindow("Test")
	w.SetContent(widget.NewVBox(widget.NewLabel("Test"), widget.NewButton("Quit", func() {
		goApp.Quit()
	})))

	w.ShowAndRun()

}
