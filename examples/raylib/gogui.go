package main

import gogui_raylib "github.com/Diegiwg/gogui/ray"

func main() {
	window := gogui_raylib.NewWindow(800, 450, "raylib backend")

	window.Init()
	defer window.Close()

	err := window.Loop()
	if err != nil {
		panic(err)
	}
}
