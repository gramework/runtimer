package runtimer

import "unsafe"

func PtrToType(ptr unsafe.Pointer) *Type {
	return (*Type)(ptr)
}
