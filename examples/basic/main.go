package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"strconv"

	gui "github.com/Diegiwg/gogui/lib"
)

func randomColor() string {
	b := make([]byte, 4)
	_, _ = rand.Read(b)
	return fmt.Sprintf("#%02x%02x%02x", b[0], b[1], b[2])
}

func main() {
	config := gui.NewConfig()
	*config.ServerAddr = "127.0.0.1"
	*config.ServerPort = 6969

	app, err := gui.NewApp(config)
	if err != nil {
		log.Fatalln(err)
	}

	main := gui.NewElement("div", "")
	main.AddChild(
		gui.NewElement("h1", "GoGui Lib Example"),
		gui.NewLabel("This is a project for learning purposes."),
	)

	interactive := gui.NewElement("div", "")

	counter := 0
	counterLabel := gui.NewLabel("Click Counter: " + strconv.Itoa(counter))
	interactive.AddChild(counterLabel, gui.NewButton("Click me!", func(widget *gui.Widget, event *gui.Event) {
		counter++
		counterLabel.SetData("content", "Click Counter: "+strconv.Itoa(counter))

		counterLabel.SetStyle("color", randomColor())
		counterLabel.Update()
	}))

	grid := gui.NewGrid(3, 3)
	for i := 0; i < 9; i++ {
		btn := gui.NewButton("Cell "+strconv.Itoa(i), func(widget *gui.Widget, event *gui.Event) {
			widget.Delete()
		})

		grid.AddChild(btn)
	}

	app.Root.AddChild(
		main,
		gui.NewElement("hr", ""),
		interactive,
		gui.NewElement("hr", ""),
		grid,
		gui.NewLabel("Click the cells to delete them!"),
		gui.NewElement("hr", ""),
	)

	err = app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
