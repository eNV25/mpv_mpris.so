package internal

import (
	"unsafe"
)

type Instance struct {
	handle *handle
}

func (ins *Instance) Run() {}

func NewInstance(handle unsafe.Pointer) *Instance {
	ins := &Instance{
		handle: handle_ref(handle),
	}
	return ins
}
