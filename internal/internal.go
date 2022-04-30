package internal

/*

#include "client.h"

void mpv_mpris_callback(void *);

static inline void mpv_mpris_set_wakeup_callback(mpv_handle *ctx, uintptr_t d) {
	mpv_set_wakeup_callback(ctx, mpv_mpris_callback, (void *)d);
}

*/
import "C"

import "unsafe"

//export mpv_mpris_callback
func mpv_mpris_callback(f func()) {
	f()
}

func mpv_set_wakeup_callback(ctx *mpv_handle, f func()) {
	C.mpv_mpris_set_wakeup_callback(ctx, func_t(f).cuintptr())
}

type func_t func()

func (f func_t) cuintptr() C.uintptr_t   { return C.uintptr_t(f.uintptr()) }
func (f func_t) uintptr() uintptr        { return *(*uintptr)(unsafe.Pointer(&f)) }
func (f func_t) pointer() unsafe.Pointer { return *(*unsafe.Pointer)(unsafe.Pointer(&f)) }
