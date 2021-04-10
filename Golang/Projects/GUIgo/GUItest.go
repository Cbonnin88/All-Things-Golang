package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Creating the Gui Window
	a := app.New()
	w := a.NewWindow("Go GUI")
	lightTheme := theme.LightTheme() // changes the color theme from dark to light
	a.Settings().SetTheme(lightTheme)

	// Creating a label for the GUI
	hello := widget.NewLabel("Watch me change")

	// creating a container for the content within the GUI
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Click here", func() { // Creates a button
			hello.SetText("This is my first GUI in Go") // Sets a text when you use the button
		}),
		widget.NewButton("Quit", func() { // Creates a quit button function
			a.Quit()
		})))

	// Resizing the GUI Window
	w.Resize(fyne.NewSize(500, 400))
	w.ShowAndRun() // running the GUI
}
