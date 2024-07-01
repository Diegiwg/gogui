package main

import (
	"log"
	"os/exec"

	gui "github.com/Diegiwg/gogui/lib"
)

func putNumber(panel *gui.Widget, num string) gui.EventHandler {
	return func(widget *gui.Widget, payload *gui.EventPayload) {
		panel.SetData("content", panel.GetData("content").(string)+num)
	}
}

func performOperation(panel *gui.Widget, op string) gui.EventHandler {
	return func(widget *gui.Widget, payload *gui.EventPayload) {
		panel.SetData("content", panel.GetData("content").(string)+op)
	}
}

func calculate(panel *gui.Widget) gui.EventHandler {
	return func(widget *gui.Widget, payload *gui.EventPayload) {
		content := panel.GetData("content").(string)

		cmd := exec.Command("python", "-c", "print("+content+")")
		output, err := cmd.Output()
		if err != nil {
			log.Fatalf("Command execution failed: %v", err)
		}

		panel.SetData("content", string(output))
	}
}

func makeCalculator() *gui.Widget {
	grid := gui.NewGrid(5, 4)
	grid.SetStyle("width", "400px")

	panel := gui.NewLabel("")
	panel.SetStyle("grid-column", "1 / span 4")
	panel.SetStyle("text-align", "right")

	grid.AddChild(
		panel,

		gui.NewButton("7", putNumber(panel, "7")), gui.NewButton("8", putNumber(panel, "8")), gui.NewButton("9", putNumber(panel, "9")), gui.NewButton("/", performOperation(panel, "/")),
		gui.NewButton("4", putNumber(panel, "4")), gui.NewButton("5", putNumber(panel, "5")), gui.NewButton("6", putNumber(panel, "6")), gui.NewButton("*", performOperation(panel, "*")),
		gui.NewButton("1", putNumber(panel, "1")), gui.NewButton("2", putNumber(panel, "2")), gui.NewButton("3", putNumber(panel, "3")), gui.NewButton("-", performOperation(panel, "-")),
		gui.NewButton("0", putNumber(panel, "0")), gui.NewButton(".", performOperation(panel, ".")), gui.NewButton("=", calculate(panel)), gui.NewButton("+", performOperation(panel, "+")),
	)

	return grid
}

func main() {
	config := gui.NewConfig()
	*config.ServerAddr = "127.0.0.1"
	*config.ServerPort = 6969

	app, err := gui.NewApp(config)
	if err != nil {
		log.Fatalln(err)
	}

	app.Root.AddChild(makeCalculator())

	err = app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
