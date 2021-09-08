package main

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"log"
	"strconv"
)

type videoGames struct {
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

	var videoGames []videoGames
	err = json.Unmarshal(games, &videoGames)
	if err != nil {
		log.Fatal(err)
	}

	// Creating the Gui Window
	a := app.New()
	w := a.NewWindow("Video Game List")
	lightTheme := theme.LightTheme() // changes the color theme from dark to light
	a.Settings().SetTheme(lightTheme)
	a.Settings().ThemeVariant()

	// Creating a list View for the GUI
	listView := widget.NewList(func() int {
		return len(videoGames)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("label template")
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		object.(*widget.Label).Text = videoGames[id].Title
	})

	// Adding a callback function that updates our content:
	contentText := widget.NewLabel("Please Select a Video Game")
	contentText.Wrapping = fyne.TextWrapWord
	listView.OnSelected = func(id widget.ListItemID) {
		contentText.TextStyle.Italic = true
		contentText.TextStyle.Bold = true
		contentText.Text = "ID number: " + strconv.Itoa(videoGames[id].ID) + "\n\n" + "Console: " + videoGames[id].Console + "\n\n" + "Original Language: " + videoGames[id].Language +
			"\n\n" + "Release Date: " + videoGames[id].Release + "\n\n" + "Rating: " + strconv.Itoa(videoGames[id].Rating) +
			"\n\n" + "Summary:\n" + videoGames[id].Overview
	}

	split := container.NewHSplit(
		listView,
		container.NewMax(contentText),
	)

	split.Offset = 0.5
	w.SetContent(split)

	// Resizing the GUI Window
	w.Resize(fyne.NewSize(1200, 900))
	w.ShowAndRun() // running the GUI
}
