package gogui_raylib

/*
#cgo CFLAGS:  -I/data/hdd1/gogui/ray/raylib-5.0_linux_amd64/include
#cgo LDFLAGS: -L/data/hdd1/gogui/ray/raylib-5.0_linux_amd64/lib -Wl,-rpath -Wl,/data/hdd1/gogui/ray/raylib-5.0_linux_amd64/lib
#cgo LDFLAGS: -L/data/hdd1/gogui/ray/raylib-5.0_linux_amd64/lib/libraylib.a -lraylib -lGL -lm -lpthread -ldl -lrt

#include <stdlib.h>
#include "raylib.h"
*/
import "C"
