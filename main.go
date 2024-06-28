package main

import (
	"github.com/Diegiwg/gogui/gogui"
)

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
	app.Button("Click me!")

	println("WARNING: This project is still in development.")
	err = app.Run()
	if err != nil {
		panic(err)
	}
}
