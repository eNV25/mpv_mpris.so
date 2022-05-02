package internal

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/env25/mpv_mpris.so/internal/util"
)

type Instance struct {
	handle *handle
	events util.Chan[*event]
	quit   util.Chan[struct{}]
}

func (ins *Instance) Run() int32 {
	ins.events.Make()
	ins.quit.Make()

	go func() {
		events := ins.events.RecieveOnly()
		for range events {
		}
	}()

	handle := ins.handle
	events := ins.events.SendOnly()
	quit := ins.quit.SendOnly()

	defer close(events)
	defer close(quit)

	fmt.Println(client_name(handle))
	set_wakeup_callback(handle, func() {
		fmt.Println(time.Now())
	})

	for {
		ev := wait_event(handle, -1)
		if ev == nil || ev.event_id == _EVENT_SHUTDOWN {
			return 0
		}
		events <- ev
	}
}

func NewInstance(handle unsafe.Pointer) *Instance {
	ins := &Instance{
		handle: handle_ref(handle),
	}
	return ins
}
