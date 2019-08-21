package main

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	// generate mine map
	generateMineMap()

	// create  window

	// create menu

	// create Container

	// add cells
	// app := app.New()

	// w := app.NewWindow("MineSweeper")
	// w.SetContent(widget.NewTabContainer(
	// 	widget.NewTabItem("game", makeGameBoard()),
	// ))

	// w.ShowAndRun()
}

// Creates a 100
func generateMineMap() [100]bool {
	var sb [100]bool

	// Create a random number generator using the time as a seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range sb {
		mine := r.Intn(100)

		if mine > 60 {
			sb[i] = true
			continue
		}

		sb[i] = false
	}

	fmt.Println(sb)
	return sb
}

func makeGameBoard() *fyne.Container {
	gridWidth := 10
	gridArea := gridWidth * gridWidth
	cells := []fyne.CanvasObject{}

	for i := 0; i < gridArea; i++ {
		cells = append(cells, makeCell())
	}

	return fyne.NewContainerWithLayout(layout.NewGridLayout(gridWidth), cells...)
}

func makeCell() *widget.Button {
	rect := widget.NewButton("M", func() { fmt.Println("yay") })
	return rect
}
