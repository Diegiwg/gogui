package main

import (
	"strconv"

	"github.com/Diegiwg/gogui/gogui"
	gogui_widgets "github.com/Diegiwg/gogui/gogui/widgets"
)

func clickHandler(counter *int, label *gogui_widgets.Label) func(ctx *gogui.HttpCtx) {
	return func(ctx *gogui.HttpCtx) {
		*counter++
		label.SetText("Click Counter: " + strconv.Itoa(*counter))
	}
}

func main() {
	config := gogui.NewConfig()
	*config.ServerAddr = "127.0.0.1"
	*config.ServerPort = 6969

	app, err := gogui.NewApp(config)
	if err != nil {
		panic(err)
	}

	app.Element("h1", "GoGui")
	app.Label("This is a project for learning purposes.")

	app.Element("hr", "")

	counter := 0
	clickLabel, _, _ := app.Label("Click Counter: " + strconv.Itoa(counter))
	app.Button("Click me!", clickHandler(&counter, clickLabel))

	println("WARNING: This project is still in development.")
	err = app.Run()
	if err != nil {
		panic(err)
	}
}
