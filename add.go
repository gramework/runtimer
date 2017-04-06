package runtimer

import "unsafe"

func Add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return add(p, x)
}

//go:linkname add runtime.add
func add(p unsafe.Pointer, x uintptr) unsafe.Pointer
