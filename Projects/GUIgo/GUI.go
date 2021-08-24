package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"log"
)

/*type GameResults struct {
	Results []VideoGames `json:"results"`
} */

type VideoGames struct {
	Title    string `json:"title"`
	ID       int    `json:"id"`
	Console  string `json:"console"`
	Language string `json:"original_lan"`
	Release  string `json:"release_date"`
	Rating   int    `json:"rating"`
	Overview string `json:"overview"`
}

func main() {

	games, err := ioutil.ReadFile("Projects/GUIgo/videoGameData.json")
	if err != nil {
		log.Fatal(err)
	}

	var videoGames []VideoGames
	err = json.Unmarshal(games, &videoGames)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v/n", videoGames)
	return

	// Creating the Gui Window
	a := app.New()
	w := a.NewWindow("Login")
	lightTheme := theme.LightTheme() // changes the color theme from dark to light
	a.Settings().SetTheme(lightTheme)
	a.Settings().ThemeVariant()

	// Creating a label for the GUI
	hello := widget.NewLabel("Welcome to Bob's bank, ")

	// creating a container for the content within the GUI
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Click Here", func() { // Creates a button
			hello.SetText("Bob's bank is an online banking app" +
				"") // Sets a text when you use the button
		}),
		widget.NewButton("About Bob's Bank", func() {
			hello.SetText("Our Mission :\n\nMaking Banking easier")
		}),
		widget.NewButton("Log out", func() { // Creates a quit button function
			a.Quit()
		})))

	// Resizing the GUI Window
	w.Resize(fyne.NewSize(500, 400))
	w.ShowAndRun() // running the GUI
}
