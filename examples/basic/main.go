package main

import (
	"log"
	"strconv"

	gui "github.com/Diegiwg/gogui/gogui"
	gui_widgets "github.com/Diegiwg/gogui/gogui/widgets"
)

func clickHandler(counter *int, label *gui_widgets.Widget) gui.HttpHandler {
	return func(ctx *gui.HttpCtx, data map[string]interface{}) {
		*counter++
		label.SetData("text", "Click Counter: "+strconv.Itoa(*counter))
	}
}

func main() {
	config := gui.NewConfig()
	*config.ServerAddr = "127.0.0.1"
	*config.ServerPort = 6969

	app, err := gui.NewApp(config)
	if err != nil {
		log.Fatalln(err)
	}

	_, title := app.Element("h1", "GoGui")
	title.Style.TextColor = "#2c43db"

	app.Label("This is a project for learning purposes.")
	_, primaryBtn := app.Button("Not click me!", nil)
	primaryBtn.SetProp("disabled", true)

	app.Element("hr", "")

	counter := 0
	_, clickLabel := app.Label("Click Counter: " + strconv.Itoa(counter))
	_, secondaryBtn := app.Button("Click me!", clickHandler(&counter, clickLabel))
	secondaryBtn.SetProp("secondary", true)

	app.Element("hr", "")

	_, grid := app.Grid(3, 3)
	for i := 1; i <= 9; i++ {
		el := grid.Element("p", "Element "+strconv.Itoa(i))
		grid.Child(el)
	}

	app.Element("hr", "")

	log.Println("WARNING: This project is still in development.")
	err = app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
