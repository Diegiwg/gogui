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

	form := gui.NewElement("div", "")

	list := gui.NewElement("ul", "")
	list.SetStyle("gap", "5px")

	inputValue := ""
	input := gui.NewInput(&inputValue)
	input.SetEvent("keyup", []string{
		"key",
		"target.value",
	}, func(widget *gui.Widget, event *gui.Event) {
		data, ok := event.Data.(map[string]interface{})
		if !ok {
			return
		}

		key := data["key"].(string)
		if key == "Enter" {
			widget.SetAttribute("value", "")
			widget.Update()

			todo := gui.NewElement("li", "")
			todo.SetStyle("display", "flex")
			todo.SetStyle("justify-content", "space-between")

			todoText := gui.NewElement("p", inputValue)
			todo.AddChild(todoText, gui.NewButton("X", func(widget *gui.Widget, event *gui.Event) {
				todo.Delete()
			}))

			list.AddChild(todo)
			list.Update()
		}

		value := data["target.value"]
		if s, ok := value.(string); ok {
			inputValue = s
		}
	})

	form.AddChild(
		gui.NewLabel("New Todo (enter for add)"),
		input,
	)

	app.Root.AddChild(
		main,
		gui.NewElement("hr", ""),
		form,
		gui.NewElement("hr", ""),
		gui.NewLabel("Todo List"),
		list,
		gui.NewElement("hr", ""),
	)

	err = app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
