package gogui_raylib

/*
#include <stdlib.h>
#include "raylib.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Window struct {
	Width  int
	Height int
	Title  string

	initialized bool
}

func NewWindow(width int, height int, title string) *Window {
	return &Window{width, height, title, false}
}

// Init initializes the raylib window
func (w *Window) Init() {
	title := C.CString(w.Title)
	defer C.free(unsafe.Pointer(title))

	C.SetTargetFPS(C.int(60))

	C.InitWindow(C.int(w.Width), C.int(w.Height), title)

	w.initialized = true
}

func (w *Window) Close() {
	C.CloseWindow()
}

func (w *Window) Loop() error {
	if !w.initialized {
		return fmt.Errorf("window not initialized")
	}

	for !C.WindowShouldClose() {
		C.BeginDrawing()
		C.ClearBackground(C.RAYWHITE)
		C.DrawText(C.CString("raylib backend"), C.int(190), C.int(200), C.int(20), C.GRAY)
		C.EndDrawing()
	}

	return nil
}
