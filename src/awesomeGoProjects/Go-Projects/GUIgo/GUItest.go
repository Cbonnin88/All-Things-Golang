package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Money Exchange")
	lightTheme := theme.LightTheme()
	a.Settings().SetTheme(lightTheme)

	hello := widget.NewLabel("Currency Converter:")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Click here to continue", func() {
			hello.SetText("Choose a currency to change over:")
		}),
		widget.NewButton("Quit", func() {
			a.Quit()
		})))
	w.Resize(fyne.NewSize(500, 400))
	w.ShowAndRun()
}
