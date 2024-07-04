package gogui_raylib

/*
#cgo CFLAGS:  -I/data/hdd1/gogui/ray/raylib-5.0_linux_amd64/include
#cgo LDFLAGS: -L/data/hdd1/gogui/ray/raylib-5.0_linux_amd64/lib -Wl,-rpath -Wl,/data/hdd1/gogui/ray/raylib-5.0_linux_amd64/lib
#cgo LDFLAGS: -lraylib -lGL -lm -lpthread -ldl -lrt

#include <stdlib.h>
#include "raylib.h"
*/
import "C"
import (
	"unsafe"
)

func Init() {
	title := C.CString("raylib backend")
	defer C.free(unsafe.Pointer(title))

	C.InitWindow(C.int(800), C.int(600), title)
	defer C.CloseWindow()

	C.SetTargetFPS(C.int(60))

	for !C.WindowShouldClose() {
		C.BeginDrawing()
		C.ClearBackground(C.RAYWHITE)
		C.DrawText(C.CString("raylib backend"), C.int(190), C.int(200), C.int(20), C.GRAY)
		C.EndDrawing()
	}
}
