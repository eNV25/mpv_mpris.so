package internal

import (
	"unsafe"
)

type Instance struct {
	handle *mpv_handle
}

func (ins *Instance) Run() int32 {
	println(mpv_client_name(ins.handle))
	return 0
}

func NewInstance(handle unsafe.Pointer) *Instance {
	ins := &Instance{
		handle: mpv_handle_ref(handle),
	}
	return ins
}
