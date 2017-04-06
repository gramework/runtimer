package runtimer

import "unsafe"

func GetEfaceDataPtr(eface interface{}) unsafe.Pointer {
	return ((*[2]unsafe.Pointer)(unsafe.Pointer(&eface))[1])
}
