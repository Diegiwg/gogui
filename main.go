package main

import "github.com/Diegiwg/gogui/gogui"

func main() {
	config := gogui.NewConfig()
	*config.ServerPort = 6969

	app, err := gogui.NewApp(config)
	if err != nil {
		panic(err)
	}

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
