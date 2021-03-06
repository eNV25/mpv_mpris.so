// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package internal

/*
#include "mpv.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

// Ref returns a reference to C object as it is.
func (x *handle) Ref() *C.mpv_handle {
	if x == nil {
		return nil
	}
	return (*C.mpv_handle)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *handle) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// handle_ref converts the C object reference into a raw struct reference without wrapping.
func handle_ref(ref unsafe.Pointer) *handle {
	return (*handle)(ref)
}

// handle_new allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func handle_new() *handle {
	return (*handle)(allocHandleMemory(1))
}

// allocHandleMemory allocates memory for type C.mpv_handle in C.
// The caller is responsible for freeing the this memory via C.free.
func allocHandleMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfHandleValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfHandleValue = unsafe.Sizeof([1]C.mpv_handle{})

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *handle) PassRef() *C.mpv_handle {
	if x == nil {
		x = (*handle)(allocHandleMemory(1))
	}
	return (*C.mpv_handle)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *node) Ref() *C.mpv_node {
	if x == nil {
		return nil
	}
	return (*C.mpv_node)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *node) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// node_ref converts the C object reference into a raw struct reference without wrapping.
func node_ref(ref unsafe.Pointer) *node {
	return (*node)(ref)
}

// node_new allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func node_new() *node {
	return (*node)(allocNodeMemory(1))
}

// allocNodeMemory allocates memory for type C.mpv_node in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNodeMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNodeValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfNodeValue = unsafe.Sizeof([1]C.mpv_node{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *node) PassRef() *C.mpv_node {
	if x == nil {
		x = (*node)(allocNodeMemory(1))
	}
	return (*C.mpv_node)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *node_list) Ref() *C.mpv_node_list {
	if x == nil {
		return nil
	}
	return (*C.mpv_node_list)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *node_list) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// node_list_ref converts the C object reference into a raw struct reference without wrapping.
func node_list_ref(ref unsafe.Pointer) *node_list {
	return (*node_list)(ref)
}

// node_list_new allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func node_list_new() *node_list {
	return (*node_list)(allocNode_listMemory(1))
}

// allocNode_listMemory allocates memory for type C.mpv_node_list in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNode_listMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNode_listValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfNode_listValue = unsafe.Sizeof([1]C.mpv_node_list{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *node_list) PassRef() *C.mpv_node_list {
	if x == nil {
		x = (*node_list)(allocNode_listMemory(1))
	}
	return (*C.mpv_node_list)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *byte_array) Ref() *C.mpv_byte_array {
	if x == nil {
		return nil
	}
	return (*C.mpv_byte_array)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *byte_array) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// byte_array_ref converts the C object reference into a raw struct reference without wrapping.
func byte_array_ref(ref unsafe.Pointer) *byte_array {
	return (*byte_array)(ref)
}

// byte_array_new allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func byte_array_new() *byte_array {
	return (*byte_array)(allocByte_arrayMemory(1))
}

// allocByte_arrayMemory allocates memory for type C.mpv_byte_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocByte_arrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfByte_arrayValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfByte_arrayValue = unsafe.Sizeof([1]C.mpv_byte_array{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *byte_array) PassRef() *C.mpv_byte_array {
	if x == nil {
		x = (*byte_array)(allocByte_arrayMemory(1))
	}
	return (*C.mpv_byte_array)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *event_property) Ref() *C.mpv_event_property {
	if x == nil {
		return nil
	}
	return (*C.mpv_event_property)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *event_property) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// event_property_ref converts the C object reference into a raw struct reference without wrapping.
func event_property_ref(ref unsafe.Pointer) *event_property {
	return (*event_property)(ref)
}

// event_property_new allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func event_property_new() *event_property {
	return (*event_property)(allocEvent_propertyMemory(1))
}

// allocEvent_propertyMemory allocates memory for type C.mpv_event_property in C.
// The caller is responsible for freeing the this memory via C.free.
func allocEvent_propertyMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfEvent_propertyValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfEvent_propertyValue = unsafe.Sizeof([1]C.mpv_event_property{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *event_property) PassRef() *C.mpv_event_property {
	if x == nil {
		x = (*event_property)(allocEvent_propertyMemory(1))
	}
	return (*C.mpv_event_property)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *event_log_message) Ref() *C.mpv_event_log_message {
	if x == nil {
		return nil
	}
	return (*C.mpv_event_log_message)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *event_log_message) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// event_log_message_ref converts the C object reference into a raw struct reference without wrapping.
func event_log_message_ref(ref unsafe.Pointer) *event_log_message {
	return (*event_log_message)(ref)
}

// event_log_message_new allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func event_log_message_new() *event_log_message {
	return (*event_log_message)(allocEvent_log_messageMemory(1))
}

// allocEvent_log_messageMemory allocates memory for type C.mpv_event_log_message in C.
// The caller is responsible for freeing the this memory via C.free.
func allocEvent_log_messageMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfEvent_log_messageValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfEvent_log_messageValue = unsafe.Sizeof([1]C.mpv_event_log_message{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *event_log_message) PassRef() *C.mpv_event_log_message {
	if x == nil {
		x = (*event_log_message)(allocEvent_log_messageMemory(1))
	}
	return (*C.mpv_event_log_message)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *event_start_file) Ref() *C.mpv_event_start_file {
	if x == nil {
		return nil
	}
	return (*C.mpv_event_start_file)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *event_start_file) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// event_start_file_ref converts the C object reference into a raw struct reference without wrapping.
func event_start_file_ref(ref unsafe.Pointer) *event_start_file {
	return (*event_start_file)(ref)
}

// event_start_file_new allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func event_start_file_new() *event_start_file {
	return (*event_start_file)(allocEvent_start_fileMemory(1))
}

// allocEvent_start_fileMemory allocates memory for type C.mpv_event_start_file in C.
// The caller is responsible for freeing the this memory via C.free.
func allocEvent_start_fileMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfEvent_start_fileValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfEvent_start_fileValue = unsafe.Sizeof([1]C.mpv_event_start_file{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *event_start_file) PassRef() *C.mpv_event_start_file {
	if x == nil {
		x = (*event_start_file)(allocEvent_start_fileMemory(1))
	}
	return (*C.mpv_event_start_file)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *event_end_file) Ref() *C.mpv_event_end_file {
	if x == nil {
		return nil
	}
	return (*C.mpv_event_end_file)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *event_end_file) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// event_end_file_ref converts the C object reference into a raw struct reference without wrapping.
func event_end_file_ref(ref unsafe.Pointer) *event_end_file {
	return (*event_end_file)(ref)
}

// event_end_file_new allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func event_end_file_new() *event_end_file {
	return (*event_end_file)(allocEvent_end_fileMemory(1))
}

// allocEvent_end_fileMemory allocates memory for type C.mpv_event_end_file in C.
// The caller is responsible for freeing the this memory via C.free.
func allocEvent_end_fileMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfEvent_end_fileValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfEvent_end_fileValue = unsafe.Sizeof([1]C.mpv_event_end_file{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *event_end_file) PassRef() *C.mpv_event_end_file {
	if x == nil {
		x = (*event_end_file)(allocEvent_end_fileMemory(1))
	}
	return (*C.mpv_event_end_file)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *event_script_input_dispatch) Ref() *C.mpv_event_script_input_dispatch {
	if x == nil {
		return nil
	}
	return (*C.mpv_event_script_input_dispatch)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *event_script_input_dispatch) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// event_script_input_dispatch_ref converts the C object reference into a raw struct reference without wrapping.
func event_script_input_dispatch_ref(ref unsafe.Pointer) *event_script_input_dispatch {
	return (*event_script_input_dispatch)(ref)
}

// event_script_input_dispatch_new allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func event_script_input_dispatch_new() *event_script_input_dispatch {
	return (*event_script_input_dispatch)(allocEvent_script_input_dispatchMemory(1))
}

// allocEvent_script_input_dispatchMemory allocates memory for type C.mpv_event_script_input_dispatch in C.
// The caller is responsible for freeing the this memory via C.free.
func allocEvent_script_input_dispatchMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfEvent_script_input_dispatchValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfEvent_script_input_dispatchValue = unsafe.Sizeof([1]C.mpv_event_script_input_dispatch{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *event_script_input_dispatch) PassRef() *C.mpv_event_script_input_dispatch {
	if x == nil {
		x = (*event_script_input_dispatch)(allocEvent_script_input_dispatchMemory(1))
	}
	return (*C.mpv_event_script_input_dispatch)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *event_client_message) Ref() *C.mpv_event_client_message {
	if x == nil {
		return nil
	}
	return (*C.mpv_event_client_message)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *event_client_message) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// event_client_message_ref converts the C object reference into a raw struct reference without wrapping.
func event_client_message_ref(ref unsafe.Pointer) *event_client_message {
	return (*event_client_message)(ref)
}

// event_client_message_new allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func event_client_message_new() *event_client_message {
	return (*event_client_message)(allocEvent_client_messageMemory(1))
}

// allocEvent_client_messageMemory allocates memory for type C.mpv_event_client_message in C.
// The caller is responsible for freeing the this memory via C.free.
func allocEvent_client_messageMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfEvent_client_messageValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfEvent_client_messageValue = unsafe.Sizeof([1]C.mpv_event_client_message{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *event_client_message) PassRef() *C.mpv_event_client_message {
	if x == nil {
		x = (*event_client_message)(allocEvent_client_messageMemory(1))
	}
	return (*C.mpv_event_client_message)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *event_hook) Ref() *C.mpv_event_hook {
	if x == nil {
		return nil
	}
	return (*C.mpv_event_hook)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *event_hook) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// event_hook_ref converts the C object reference into a raw struct reference without wrapping.
func event_hook_ref(ref unsafe.Pointer) *event_hook {
	return (*event_hook)(ref)
}

// event_hook_new allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func event_hook_new() *event_hook {
	return (*event_hook)(allocEvent_hookMemory(1))
}

// allocEvent_hookMemory allocates memory for type C.mpv_event_hook in C.
// The caller is responsible for freeing the this memory via C.free.
func allocEvent_hookMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfEvent_hookValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfEvent_hookValue = unsafe.Sizeof([1]C.mpv_event_hook{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *event_hook) PassRef() *C.mpv_event_hook {
	if x == nil {
		x = (*event_hook)(allocEvent_hookMemory(1))
	}
	return (*C.mpv_event_hook)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *event_command) Ref() *C.mpv_event_command {
	if x == nil {
		return nil
	}
	return (*C.mpv_event_command)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *event_command) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// event_command_ref converts the C object reference into a raw struct reference without wrapping.
func event_command_ref(ref unsafe.Pointer) *event_command {
	return (*event_command)(ref)
}

// event_command_new allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func event_command_new() *event_command {
	return (*event_command)(allocEvent_commandMemory(1))
}

// allocEvent_commandMemory allocates memory for type C.mpv_event_command in C.
// The caller is responsible for freeing the this memory via C.free.
func allocEvent_commandMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfEvent_commandValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfEvent_commandValue = unsafe.Sizeof([1]C.mpv_event_command{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *event_command) PassRef() *C.mpv_event_command {
	if x == nil {
		x = (*event_command)(allocEvent_commandMemory(1))
	}
	return (*C.mpv_event_command)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *event) Ref() *C.mpv_event {
	if x == nil {
		return nil
	}
	return (*C.mpv_event)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *event) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// event_ref converts the C object reference into a raw struct reference without wrapping.
func event_ref(ref unsafe.Pointer) *event {
	return (*event)(ref)
}

// event_new allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func event_new() *event {
	return (*event)(allocEventMemory(1))
}

// allocEventMemory allocates memory for type C.mpv_event in C.
// The caller is responsible for freeing the this memory via C.free.
func allocEventMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfEventValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfEventValue = unsafe.Sizeof([1]C.mpv_event{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *event) PassRef() *C.mpv_event {
	if x == nil {
		x = (*event)(allocEventMemory(1))
	}
	return (*C.mpv_event)(unsafe.Pointer(x))
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = unsafe.Pointer(p)
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - uintptr(h.Data))
	}
	return
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// rawString reperesents a string backed by data on the C side.
type rawString string

// Copy returns a Go-managed copy of raw string.
func (raw rawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(h.Data), C.int(h.Len))
}

// unpackPCharString copies the data from Go string as *C.char.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	mem0 := unsafe.Pointer(C.CString(str))
	allocs.Add(mem0)
	return (*C.char)(mem0), allocs
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// allocPCharMemory allocates memory for type *C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPCharValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfPCharValue = unsafe.Sizeof([1]*C.char{})

const sizeOfPtr = unsafe.Sizeof(&struct{}{})

// unpackArgSString transforms a sliced Go data structure into plain C format.
func unpackArgSString(x []string) (unpacked **C.char, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	len0 := len(x)
	mem0 := allocPCharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: mem0,
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.char)(unsafe.Pointer(h0))
	for i0 := range x {
		v0[i0], _ = unpackPCharString(x[i0])
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.char)(h.Data)
	return
}

// packSString reads sliced Go data structure out from plain C format.
func packSString(v []string, ptr0 **C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.char)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = packPCharString(ptr1)
	}
}

// copyPNodeBytes copies the data from Go slice as *C.mpv_node.
func copyPNodeBytes(slice *sliceHeader) (*C.mpv_node, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	mem0 := unsafe.Pointer(C.CBytes(*(*[]byte)(unsafe.Pointer(&sliceHeader{
		Data: slice.Data,
		Len:  int(sizeOfNodeValue) * slice.Len,
		Cap:  int(sizeOfNodeValue) * slice.Len,
	}))))
	allocs.Add(mem0)

	return (*C.mpv_node)(mem0), allocs
}
