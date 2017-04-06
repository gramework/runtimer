// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Runtime type representation.

package runtimer

import "unsafe"

// Tflag is documented in reflect/type.go.
//
// Tflag values must be kept in sync with copies in:
//	cmd/compile/internal/gc/reflect.go
//	cmd/link/internal/ld/decodesym.go
//	reflect/type.go
type Tflag uint8

const (
	TflagUncommon  Tflag = 1 << 0
	TflagExtraStar Tflag = 1 << 1
	TflagNamed     Tflag = 1 << 2
)

// Needs to be in sync with ../cmd/compile/internal/ld/decodesym.go:/^func.commonsize,
// ../cmd/compile/internal/gc/reflect.go:/^func.dcommontype and
// ../reflect/type.go:/^type.rtype.
type Type struct {
	Size       uintptr
	Ptrdata    uintptr // size of memory prefix holding all pointers
	Hash       uint32
	Tflag      Tflag
	Align      uint8
	Fieldalign uint8
	Kind       uint8
	Alg        *TypeAlg
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	Gcdata    *byte
	Str       NameOff
	PtrToThis TypeOff
}

func (t *Type) string() string {
	s := t.NameOff(t.Str).Name()
	if t.Tflag&TflagExtraStar != 0 {
		return s[1:]
	}
	return s
}

func (t *Type) uncommon() *uncommontype {
	if t.Tflag&TflagUncommon == 0 {
		return nil
	}
	switch t.Kind & KindMask {
	case KindStruct:
		type u struct {
			structtype
			u uncommontype
		}
		return &(*u)(unsafe.Pointer(t)).u
	case KindPtr:
		type u struct {
			ptrtype
			u uncommontype
		}
		return &(*u)(unsafe.Pointer(t)).u
	case KindFunc:
		type u struct {
			functype
			u uncommontype
		}
		return &(*u)(unsafe.Pointer(t)).u
	case KindSlice:
		type u struct {
			slicetype
			u uncommontype
		}
		return &(*u)(unsafe.Pointer(t)).u
	case KindArray:
		type u struct {
			arraytype
			u uncommontype
		}
		return &(*u)(unsafe.Pointer(t)).u
	case KindChan:
		type u struct {
			chantype
			u uncommontype
		}
		return &(*u)(unsafe.Pointer(t)).u
	case KindMap:
		type u struct {
			Maptype
			u uncommontype
		}
		return &(*u)(unsafe.Pointer(t)).u
	case KindInterface:
		type u struct {
			interfacetype
			u uncommontype
		}
		return &(*u)(unsafe.Pointer(t)).u
	default:
		type u struct {
			Type
			u uncommontype
		}
		return &(*u)(unsafe.Pointer(t)).u
	}
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func (t *Type) Name() string {
	if t.Tflag&TflagNamed == 0 {
		return ""
	}
	s := t.string()
	i := len(s) - 1
	for i >= 0 {
		if s[i] == '.' {
			break
		}
		i--
	}
	return s[i+1:]
}

// Mutual exclusion locks.  In the uncontended case,
// as fast as spin locks (just a few user-level instructions),
// but on the contention path they sleep in the kernel.
// A zeroed Mutex is unlocked (no need to initialize each lock).
type Mutex struct {
	// Futex-based impl treats it as uint32 key,
	// while sema-based impl as M* waitm.
	// Used to be a union, but unions break precise GC.
	Key uintptr
}

func ResolveNameOff(ptrInModule unsafe.Pointer, off NameOff) Name {
	return resolveNameOff(ptrInModule, off)
}

//go:linkname resolveNameOff runtime.resolveNameOff
func resolveNameOff(ptrInModule unsafe.Pointer, off NameOff) Name

func (t *Type) NameOff(off NameOff) Name {
	return resolveNameOff(unsafe.Pointer(t), off)
}

func ResolveTypeOff(ptrInModule unsafe.Pointer, off TypeOff) *Type {
	return resolveTypeOff(ptrInModule, off)
}

//go:linkname resolveTypeOff runtime.resolveTypeOff
func resolveTypeOff(ptrInModule unsafe.Pointer, off TypeOff) *Type

func (t *Type) TypeOff(off TypeOff) *Type {
	return ResolveTypeOff(unsafe.Pointer(t), off)
}

func (t *Type) TextOff(off textOff) unsafe.Pointer {
	return t.textOff(off)
}

//go:linkname Type.textOff runtime._type.textOff
func (t *Type) textOff(off textOff) unsafe.Pointer

func (t *functype) in() []*Type {
	// See funcType in reflect/type.go for details on data layout.
	uadd := uintptr(unsafe.Sizeof(functype{}))
	if t.typ.Tflag&TflagUncommon != 0 {
		uadd += unsafe.Sizeof(uncommontype{})
	}
	return (*[1 << 20]*Type)(Add(unsafe.Pointer(t), uadd))[:t.inCount]
}

func (t *functype) out() []*Type {
	// See funcType in reflect/type.go for details on data layout.
	uadd := uintptr(unsafe.Sizeof(functype{}))
	if t.typ.Tflag&TflagUncommon != 0 {
		uadd += unsafe.Sizeof(uncommontype{})
	}
	outCount := t.outCount & (1<<15 - 1)
	return (*[1 << 20]*Type)(Add(unsafe.Pointer(t), uadd))[t.inCount : t.inCount+outCount]
}

func (t *functype) dotdotdot() bool {
	return t.outCount&(1<<15) != 0
}

type NameOff int32
type TypeOff int32
type textOff int32

type method struct {
	Name NameOff
	mtyp TypeOff
	ifn  textOff
	tfn  textOff
}

type uncommontype struct {
	pkgpath NameOff
	mcount  uint16 // number of methods
	_       uint16 // unused
	moff    uint32 // offset from this uncommontype to [mcount]method
	_       uint32 // unused
}

type imethod struct {
	Name NameOff
	ityp TypeOff
}

type interfacetype struct {
	typ     Type
	pkgpath Name
	mhdr    []imethod
}

type Maptype struct {
	Typ           Type
	Key           *Type
	Elem          *Type
	Bucket        *Type  // internal type representing a hash bucket
	Hmap          *Type  // internal type representing a hmap
	Keysize       uint8  // size of key slot
	Indirectkey   bool   // store ptr to key instead of key itself
	Valuesize     uint8  // size of value slot
	Indirectvalue bool   // store ptr to value instead of value itself
	Bucketsize    uint16 // size of bucket
	Reflexivekey  bool   // true if k==k for all keys
	Needkeyupdate bool   // true if we need to update key on an overwrite
}

type arraytype struct {
	typ   Type
	elem  *Type
	slice *Type
	len   uintptr
}

type chantype struct {
	typ  Type
	elem *Type
	dir  uintptr
}

type slicetype struct {
	typ  Type
	elem *Type
}

type functype struct {
	typ      Type
	inCount  uint16
	outCount uint16
}

type ptrtype struct {
	typ  Type
	elem *Type
}

type structfield struct {
	Name       Name
	typ        *Type
	offsetAnon uintptr
}

func (f *structfield) offset() uintptr {
	return f.offsetAnon >> 1
}

type structtype struct {
	typ     Type
	pkgPath Name
	fields  []structfield
}

// Name is an encoded type Name with optional extra data.
// See reflect/type.go for details.
type Name struct {
	Bytes *byte
}

func (n Name) data(off int) *byte {
	return (*byte)(Add(unsafe.Pointer(n.Bytes), uintptr(off)))
}

func (n Name) isExported() bool {
	return (*n.Bytes)&(1<<0) != 0
}

func (n Name) NameLen() int {
	return int(uint16(*n.data(1))<<8 | uint16(*n.data(2)))
}

func (n Name) tagLen() int {
	if *n.data(0)&(1<<1) == 0 {
		return 0
	}
	off := 3 + n.NameLen()
	return int(uint16(*n.data(off))<<8 | uint16(*n.data(off + 1)))
}

func (n Name) Name() (s string) {
	if n.Bytes == nil {
		return ""
	}
	nl := n.NameLen()
	if nl == 0 {
		return ""
	}
	hdr := (*StringStruct)(unsafe.Pointer(&s))
	hdr.Str = unsafe.Pointer(n.data(3))
	hdr.Len = nl
	return s
}

func (n Name) tag() (s string) {
	tl := n.tagLen()
	if tl == 0 {
		return ""
	}
	nl := n.NameLen()
	hdr := (*StringStruct)(unsafe.Pointer(&s))
	hdr.Str = unsafe.Pointer(n.data(3 + nl + 2))
	hdr.Len = tl
	return s
}

func (n Name) pkgPath() string {
	if n.Bytes == nil || *n.data(0)&(1<<2) == 0 {
		return ""
	}
	off := 3 + n.NameLen()
	if tl := n.tagLen(); tl > 0 {
		off += 2 + tl
	}
	var NameOff NameOff
	copy((*[4]byte)(unsafe.Pointer(&NameOff))[:], (*[4]byte)(unsafe.Pointer(n.data(off)))[:])
	pkgPathName := resolveNameOff(unsafe.Pointer(n.Bytes), NameOff)
	return pkgPathName.Name()
}

// TypesEqual reports whether two types are equal.
//
// Everywhere in the runtime and reflect packages, it is assumed that
// there is exactly one *Type per Go type, so that pointer equality
// can be used to test if types are equal. There is one place that
// breaks this assumption: buildmode=shared. In this case a type can
// appear as two different pieces of memory. This is hidden from the
// runtime and reflect package by the per-module typemap built in
// typelinksinit. It uses typesEqual to map types from later modules
// back into earlier ones.
//
// Only typelinksinit needs this function.
func TypesEqual(t, v *Type) bool {
	return typesEqual(t, v)
}

//go:linkname typesEqual runtime.typesEqual
func typesEqual(t, v *Type) bool
