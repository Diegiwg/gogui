package main

import (
	"log"

	gui "github.com/Diegiwg/gogui/lib"
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

	app.Root.AddChild(
		main,
		gui.NewElement("hr", ""),
		interactive,
		gui.NewElement("hr", ""),
	)

	err = app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
