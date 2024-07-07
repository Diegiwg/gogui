package main

import (
	"context"
	"log"
	"strconv"

	gui "github.com/Diegiwg/gogui/lib"
	"nhooyr.io/websocket"
)

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
	interactive.AddChild(counterLabel, gui.NewButton("Click me!", func(widget *gui.Widget, event *gui.Event, conn *websocket.Conn, ctx *context.Context) {
		counter++
		counterLabel.SetData("content", "Click Counter: "+strconv.Itoa(counter))
	}))

	grid := gui.NewGrid(3, 3)
	for i := 0; i < 9; i++ {
		grid.AddChild(gui.NewButton("Cell "+strconv.Itoa(i), nil))
	}
	grid.AddChild(gui.NewLabel("Click the cells to delete them!"))

	app.Root.AddChild(
		main,
		gui.NewElement("hr", ""),
		interactive,
		gui.NewElement("hr", ""),
		grid,
		gui.NewElement("hr", ""),
	)

	err = app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
