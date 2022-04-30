package main

import "C"

import (
	"unsafe"

	"github.com/env25/mpv_mpris.so/internal"
)

func main() {}

//export mpv_open_cplugin
func mpv_open_cplugin(handle unsafe.Pointer) C.int {
	internal.NewInstance(handle).Run()
	return 0
}
