package main

import (
	"strconv"

	"github.com/Diegiwg/gogui/gogui"
	gogui_widgets "github.com/Diegiwg/gogui/gogui/widgets"
)

func clickHandler(counter *int, label *gogui_widgets.Widget) func(ctx *gogui.HttpCtx, data map[string]interface{}) {
	return func(ctx *gogui.HttpCtx, data map[string]interface{}) {
		*counter++
		label.SetData("text", "Click Counter: "+strconv.Itoa(*counter))
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
	_, clickLabel := app.Label("Click Counter: " + strconv.Itoa(counter))
	app.Button("Click me!", clickHandler(&counter, clickLabel))

	_, grid := app.Grid(3, 3)
	for i := 1; i <= 9; i++ {
		el := grid.Element("p", strconv.Itoa(i))
		grid.Child(el)
	}

	println("WARNING: This project is still in development.")
	err = app.Run()
	if err != nil {
		panic(err)
	}
}
