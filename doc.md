# runtimer
--
    import "github.com/ninedraft/runtimer"

Package runtimer provides you unsafe way to use runtime internals

## Usage

```go
const (
	C0 = uintptr((8-PtrSize)/4*2860486313 + (PtrSize-4)/4*33054211828000289)
	C1 = uintptr((8-PtrSize)/4*3267000013 + (PtrSize-4)/4*23344194077549503)
)
```

```go
const (
	AlgNOEQ = iota
	AlgMEM0
	AlgMEM8
	AlgMEM16
	AlgMEM32
	AlgMEM64
	AlgMEM128
	AlgSTRING
	AlgINTER
	AlgNILINTER
	AlgFLOAT32
	AlgFLOAT64
	AlgCPLX64
	AlgCPLX128
	AlgMax
)
```
type algorithms - known to compiler

```go
const (
	KindBool = 1 + iota
	KindInt
	KindInt8
	KindInt16
	KindInt32
	KindInt64
	KindUint
	KindUint8
	KindUint16
	KindUint32
	KindUint64
	KindUintptr
	KindFloat32
	KindFloat64
	KindComplex64
	KindComplex128
	KindArray
	KindChan
	KindFunc
	KindInterface
	KindMap
	KindPtr
	KindSlice
	KindString
	KindStruct
	KindUnsafePointer

	KindDirectIface = 1 << 5
	KindGCProg      = 1 << 6
	KindNoPointers  = 1 << 7
	KindMask        = (1 << 5) - 1
)
```

```go
const (
	//go:linkname GOOS runtime.internal.sys.GOOS
	GOOS = `unknown`

	//go:linkname GoosAndroid runtime.internal.sys.GoosAndroid
	GoosAndroid uint = 0
	//go:linkname GoosDarwin runtime.internal.sys.GoosDarwin
	GoosDarwin uint = 0
	//go:linkname GoosDragonfly runtime.internal.sys.GoosDragonfly
	GoosDragonfly uint = 0
	//go:linkname GoosFreebsd runtime.internal.sys.GoosFreebsd
	GoosFreebsd uint = 0
	//go:linkname GoosLinux runtime.internal.sys.GoosLinux
	GoosLinux uint = 0
	//go:linkname GoosNacl runtime.internal.sys.GoosNacl
	GoosNacl uint = 0
	//go:linkname GoosNetbsd runtime.internal.sys.GoosNetbsd
	GoosNetbsd uint = 0
	//go:linkname GoosOpenbsd runtime.internal.sys.GoosOpenbsd
	GoosOpenbsd uint = 0
	//go:linkname GoosPlan9 runtime.internal.sys.GoosPlan9
	GoosPlan9 uint = 0
	//go:linkname GoosSolaris runtime.internal.sys.GoosSolaris
	GoosSolaris uint = 0
	//go:linkname GoosWindows runtime.internal.sys.GoosWindows
	GoosWindows uint = 0

	//go:linkname Goarch386 runtime.internal.sys.Goarch386
	Goarch386 uint = 0
	//go:linkname GoarchAmd64 runtime.internal.sys.GoarchAmd64
	GoarchAmd64 uint = 0
	//go:linkname GoarchAmd64p32 runtime.internal.sys.GoarchAmd64p32
	GoarchAmd64p32 uint = 0
	//go:linkname GoarchArm runtime.internal.sys.GoarchArm
	GoarchArm uint = 0
	//go:linkname GoarchArmbe runtime.internal.sys.GoarchArmbe
	GoarchArmbe uint = 0
	//go:linkname GoarchArm64 runtime.internal.sys.GoarchArm64
	GoarchArm64 uint = 0
	//go:linkname GoarchArm64be runtime.internal.sys.GoarchArm64be
	GoarchArm64be uint = 0
	//go:linkname GoarchPpc64 runtime.internal.sys.GoarchPpc64
	GoarchPpc64 uint = 0
	//go:linkname GoarchPpc64le runtime.internal.sys.GoarchPpc64le
	GoarchPpc64le uint = 0
	//go:linkname GoarchMips runtime.internal.sys.GoarchMips
	GoarchMips uint = 0
	//go:linkname GoarchMipsle runtime.internal.sys.GoarchMipsle
	GoarchMipsle uint = 0
	//go:linkname GoarchMips64 runtime.internal.sys.GoarchMips64
	GoarchMips64 uint = 0
	//go:linkname GoarchMips64le runtime.internal.sys.GoarchMips64le
	GoarchMips64le uint = 0
	//go:linkname GoarchMips64p32 runtime.internal.sys.GoarchMips64p32
	GoarchMips64p32 uint = 0
	//go:linkname GoarchMips64p32le runtime.internal.sys.GoarchMips64p32le
	GoarchMips64p32le uint = 0
	//go:linkname GoarchPpc runtime.internal.sys.GoarchPpc
	GoarchPpc uint = 0
	//go:linkname GoarchS390 runtime.internal.sys.GoarchS390
	GoarchS390 uint = 0
	//go:linkname GoarchS390x runtime.internal.sys.GoarchS390x
	GoarchS390x uint = 0
	//go:linkname GoarchSparc runtime.internal.sys.GoarchSparc
	GoarchSparc uint = 0
	//go:linkname GoarchSparc64 runtime.internal.sys.GoarchSparc64
	GoarchSparc64 uint = 0
)
```

```go
const (
	PageShift uint = 13

	// Public64bit = 1 on 64-bit systems, 0 on 32-bit systems
	Public64bit uint = 1 << (^uintptr(0) >> 63) / 2

	MHeapMapTotalBits = (Public64bit*GoosWindows)*35 + (Public64bit*(1-GoosWindows)*(1-GoosDarwin*GoarchArm64))*39 + GoosDarwin*GoarchArm64*31 + (1-Public64bit)*(32-(GoarchMips+GoarchMipsle))
	MHeapMapBits      = MHeapMapTotalBits - PageShift

	// MaxMem is the maximum heap arena size minus 1.
	//
	// On 32-bit, this is also the maximum heap pointer value,
	// since the arena starts at address 0.
	MaxMem = 1<<MHeapMapTotalBits - 1
)
```

```go
const HashRandomBytes = PtrSize / 4 * 64
```

```go
const Msanenabled = false
```

```go
const Msanenabled = true
```

```go
const PtrSize = 4 << (^uintptr(0) >> 63) // unsafe.Sizeof(uintptr(0)) but an ideal const

```
go:linkname PtrSize runtime.internal.sys.PtrSize

```go
var AlgArray = [AlgMax]TypeAlg{
	AlgNOEQ:     {nil, nil},
	AlgMEM0:     {Memhash0, Memequal0},
	AlgMEM8:     {Memhash8, Memequal8},
	AlgMEM16:    {Memhash16, Memequal16},
	AlgMEM32:    {Memhash32, Memequal32},
	AlgMEM64:    {Memhash64, Memequal64},
	AlgMEM128:   {Memhash128, Memequal128},
	AlgSTRING:   {Strhash, Strequal},
	AlgINTER:    {Interhash, Interequal},
	AlgNILINTER: {Nilinterhash, Nilinterequal},
	AlgFLOAT32:  {F32hash, F32equal},
	AlgFLOAT64:  {F64hash, F64equal},
	AlgCPLX64:   {C64hash, C64equal},
	AlgCPLX128:  {C128hash, C128equal},
}
```

```go
var CPUIDECX uint32
```
go:linkname CPUIDECX runtime.cpuid_ecx

#### func  Add

```go
func Add(p unsafe.Pointer, x uintptr) unsafe.Pointer
```

#### func  Aeshash

```go
func Aeshash(p unsafe.Pointer, h, s uintptr) uintptr
```

#### func  Aeshash32

```go
func Aeshash32(p unsafe.Pointer, h uintptr) uintptr
```

#### func  Aeshash64

```go
func Aeshash64(p unsafe.Pointer, h uintptr) uintptr
```

#### func  Aeshashstr

```go
func Aeshashstr(p unsafe.Pointer, h uintptr) uintptr
```

#### func  Atoi

```go
func Atoi(s string) (int, bool)
```
Atoi parses an int from a string s. The bool result reports whether s is a
number representable by a value of type int.

#### func  Atoi32

```go
func Atoi32(s string) (int32, bool)
```
Atoi32 is like Atoi but for integers that fit into an int32.

#### func  BytesHash

```go
func BytesHash(b []byte, seed uintptr) uintptr
```

#### func  C128equal

```go
func C128equal(p, q unsafe.Pointer) bool
```

#### func  C128hash

```go
func C128hash(p unsafe.Pointer, h uintptr) uintptr
```

#### func  C64equal

```go
func C64equal(p, q unsafe.Pointer) bool
```

#### func  C64hash

```go
func C64hash(p unsafe.Pointer, h uintptr) uintptr
```

#### func  Cas

```go
func Cas(ptr *uint32, old, new uint32) bool
```

#### func  Casp1

```go
func Casp1(ptr *unsafe.Pointer, old, new unsafe.Pointer) bool
```

#### func  Casuintptr

```go
func Casuintptr(ptr *uintptr, old, new uintptr) bool
```

#### func  Contains

```go
func Contains(s, t string) bool
```

#### func  EfaceDataPtr

```go
func EfaceDataPtr(eface interface{}) *unsafe.Pointer
```

#### func  EfaceHash

```go
func EfaceHash(i interface{}, seed uintptr) uintptr
```

#### func  Efaceeq

```go
func Efaceeq(t *Type, x, y unsafe.Pointer) bool
```

#### func  Encoderune

```go
func Encoderune(p []byte, r rune) int
```

#### func  F32equal

```go
func F32equal(p, q unsafe.Pointer) bool
```

#### func  F32hash

```go
func F32hash(p unsafe.Pointer, h uintptr) uintptr
```

#### func  F64equal

```go
func F64equal(p, q unsafe.Pointer) bool
```

#### func  F64hash

```go
func F64hash(p unsafe.Pointer, h uintptr) uintptr
```

#### func  Fastrand

```go
func Fastrand() uint32
```

#### func  Findnull

```go
func Findnull(s *byte) int
```

#### func  Findnullw

```go
func Findnullw(s *uint16) int
```

#### func  GetEfaceDataPtr

```go
func GetEfaceDataPtr(eface interface{}) unsafe.Pointer
```

#### func  Gostringnocopy

```go
func Gostringnocopy(str *byte) string
```

#### func  HasPrefix

```go
func HasPrefix(s, t string) bool
```

#### func  IfaceHash

```go
func IfaceHash(i interface {
	F()
}, seed uintptr) uintptr
```

#### func  Ifaceeq

```go
func Ifaceeq(t *Itab, x, y unsafe.Pointer) bool
```

#### func  Index

```go
func Index(s, t string) int
```

#### func  Int32Hash

```go
func Int32Hash(i uint32, seed uintptr) uintptr
```

#### func  Int64Hash

```go
func Int64Hash(i uint64, seed uintptr) uintptr
```

#### func  Interequal

```go
func Interequal(p, q unsafe.Pointer) bool
```

#### func  Interhash

```go
func Interhash(p unsafe.Pointer, h uintptr) uintptr
```

#### func  IsDirectIface

```go
func IsDirectIface(t *Type) bool
```
IsDirectIface reports whether t is stored directly in an interface value.

#### func  Loadint64

```go
func Loadint64(ptr *int64) int64
```

#### func  Loaduint

```go
func Loaduint(ptr *uint) uint
```

#### func  Loaduintptr

```go
func Loaduintptr(ptr *uintptr) uintptr
```

#### func  Lock

```go
func Lock(l *Mutex)
```

#### func  MemclrHasPointers

```go
func MemclrHasPointers(ptr unsafe.Pointer, n uintptr)
```

#### func  MemclrNoHeapPointers

```go
func MemclrNoHeapPointers(ptr unsafe.Pointer, n uintptr)
```

#### func  Memequal

```go
func Memequal(a, b unsafe.Pointer, size uintptr) bool
```

#### func  Memequal0

```go
func Memequal0(p, q unsafe.Pointer) bool
```

#### func  Memequal128

```go
func Memequal128(p, q unsafe.Pointer) bool
```

#### func  Memequal16

```go
func Memequal16(p, q unsafe.Pointer) bool
```

#### func  Memequal32

```go
func Memequal32(p, q unsafe.Pointer) bool
```

#### func  Memequal64

```go
func Memequal64(p, q unsafe.Pointer) bool
```

#### func  Memequal8

```go
func Memequal8(p, q unsafe.Pointer) bool
```

#### func  Memhash0

```go
func Memhash0(p unsafe.Pointer, h uintptr) uintptr
```

#### func  Memhash128

```go
func Memhash128(p unsafe.Pointer, h uintptr) uintptr
```

#### func  Memhash16

```go
func Memhash16(p unsafe.Pointer, h uintptr) uintptr
```

#### func  Memhash32

```go
func Memhash32(p unsafe.Pointer, h uintptr) uintptr
```

#### func  Memhash64

```go
func Memhash64(p unsafe.Pointer, h uintptr) uintptr
```

#### func  Memhash8

```go
func Memhash8(p unsafe.Pointer, h uintptr) uintptr
```

#### func  MemhashVarlen

```go
func MemhashVarlen(p unsafe.Pointer, h uintptr) uintptr
```
MemhashVarlen is defined in runtime assembly because it needs access to the
closure. It appears here to provide an argument signature for the assembly
routine.

#### func  Msanread

```go
func Msanread(addr unsafe.Pointer, sz uintptr)
```

#### func  Newarray

```go
func Newarray(typ *Type, n int) unsafe.Pointer
```

#### func  Newobject

```go
func Newobject(typ *Type) unsafe.Pointer
```

#### func  Nilinterequal

```go
func Nilinterequal(p, q unsafe.Pointer) bool
```

#### func  Nilinterhash

```go
func Nilinterhash(p unsafe.Pointer, h uintptr) uintptr
```

#### func  Noescape

```go
func Noescape(p unsafe.Pointer) unsafe.Pointer
```

#### func  Or8

```go
func Or8(ptr *uint8, val uint8)
```

#### func  PtrPtrToStringPtr

```go
func PtrPtrToStringPtr(ptr *unsafe.Pointer) *string
```

#### func  PtrToString

```go
func PtrToString(ptr unsafe.Pointer) string
```

#### func  PtrToStringPtr

```go
func PtrToStringPtr(ptr unsafe.Pointer) *string
```

#### func  Rawstring

```go
func Rawstring(size int) (s string, b []byte)
```

#### func  Storeuintptr

```go
func Storeuintptr(ptr *uintptr, new uintptr)
```

#### func  Strequal

```go
func Strequal(p, q unsafe.Pointer) bool
```

#### func  Strhash

```go
func Strhash(p unsafe.Pointer, h uintptr) uintptr
```

#### func  StringHash

```go
func StringHash(s string, seed uintptr) uintptr
```

#### func  Throw

```go
func Throw(s string)
```

#### func  Typedmemclr

```go
func Typedmemclr(typ *Type, ptr unsafe.Pointer)
```

#### func  Typedmemmove

```go
func Typedmemmove(typ *Type, dst, src unsafe.Pointer)
```

#### func  TypesEqual

```go
func TypesEqual(t, v *Type) bool
```
TypesEqual reports whether two Types are equal.

Everywhere in the runtime and reflect packages, it is assumed that there is
exactly one *Type per Go Type, so that pointer equality can be used to test if
Types are equal. There is one place that breaks this assumption:
buildmode=shared. In this case a Type can appear as two different pieces of
memory. This is hidden from the runtime and reflect package by the per-module
Typemap built in Typelinksinit. It uses TypesEqual to map Types from later
modules back into earlier ones.

Only Typelinksinit needs this function.

#### func  Unlock

```go
func Unlock(l *Mutex)
```

#### func  Xaddint64

```go
func Xaddint64(ptr *int64, delta int64) int64
```

#### type ArrayType

```go
type ArrayType struct {
	Typ   Type
	Elem  *Type
	Slice *Type
	Len   uintptr
}
```


#### type ChanType

```go
type ChanType struct {
	Typ  Type
	Elem *Type
	Dir  uintptr
}
```


#### type FuncType

```go
type FuncType struct {
}
```


#### func (*FuncType) Dotdotdot

```go
func (t *FuncType) Dotdotdot() bool
```

#### type Hex

```go
type Hex uint64
```

go:linkname Hex runtime.hex

#### type InterfaceType

```go
type InterfaceType struct {
}
```


#### type Itab

```go
type Itab struct {
	Inter  *InterfaceType
	Type   *Type
	Link   *Itab
	Hash   uint32 // copy of _type.hash. Used for type switches.
	Bad    bool   // type does not implement interface
	Inhash bool   // has this itab been added to hash?
	Unused [2]byte
	Fun    [1]uintptr // variable sized
}
```

layout of Itab known to compilers allocated in non-garbage-collected memory
Needs to be in sync with
../cmd/compile/internal/gc/reflect.go:/^func.dumptypestructs.

#### type MapType

```go
type MapType struct {
	Typ           Type
	Key           *Type
	Elem          *Type
	Bucket        *Type  // internal Type representing a hash bucket
	Hmap          *Type  // internal Type representing a hmap
	Keysize       uint8  // size of key slot
	Indirectkey   bool   // store ptr to key instead of key itself
	Valuesize     uint8  // size of value slot
	Indirectvalue bool   // store ptr to value instead of value itself
	Bucketsize    uint16 // size of bucket
	Reflexivekey  bool   // true if k==k for all keys
	Needkeyupdate bool   // true if we need to update key on an overwrite
}
```


#### type Mutex

```go
type Mutex struct {
	// Futex-based impl treats it as uint32 key,
	// while sema-based impl as M* waitm.
	// Used to be a union, but unions break precise GC.
	Key uintptr
}
```

Mutex - Mutual exclusion locks. In the uncontended case, as fast as spin locks
(just a few user-level instructions), but on the contention path they sleep in
the kernel. A zeroed Mutex is unlocked (no need to initialize each lock).

#### type Name

```go
type Name struct {
	Bytes *byte
}
```

Name is an encoded Type Name with optional extra data. See reflect/Type.go for
details.

#### func  ResolveNameOff

```go
func ResolveNameOff(ptrInModule unsafe.Pointer, off NameOff) Name
```

#### func (Name) Data

```go
func (n Name) Data(off int) *byte
```

#### func (Name) IsExported

```go
func (n Name) IsExported() bool
```

#### func (Name) Name

```go
func (n Name) Name() (s string)
```

#### func (Name) NameLen

```go
func (n Name) NameLen() int
```

#### func (Name) PkgPath

```go
func (n Name) PkgPath() string
```

#### func (Name) Tag

```go
func (n Name) Tag() (s string)
```

#### func (Name) TagLen

```go
func (n Name) TagLen() int
```

#### type NameOff

```go
type NameOff int32
```


#### type PtrType

```go
type PtrType struct {
}
```


#### type SliceType

```go
type SliceType struct {
	Typ  Type
	Elem *Type
}
```


#### type StringStruct

```go
type StringStruct struct {
	Str unsafe.Pointer
	Len int
}
```


#### func  StringStructOf

```go
func StringStructOf(sp *string) *StringStruct
```

#### type StringStructDWARF

```go
type StringStructDWARF struct {
	Str *byte
	Len int
}
```

Variant with *byte pointer type for DWARF debugging.

#### type StructType

```go
type StructType struct {
	Typ     Type
	PkgPath Name
	Fields  []Structfield
}
```


#### type Structfield

```go
type Structfield struct {
	Name       Name
	Typ        *Type
	OffsetAnon uintptr
}
```


#### func (*Structfield) Offset

```go
func (f *Structfield) Offset() uintptr
```

#### type Tflag

```go
type Tflag uint8
```

Tflag is documented in reflect/Type.go.

Tflag values must be kept in sync with copies in:

    cmd/compile/internal/gc/reflect.go
    cmd/link/internal/ld/decodesym.go
    reflect/Type.go

```go
const (
	TflagUncommon  Tflag = 1 << 0
	TflagExtraStar Tflag = 1 << 1
	TflagNamed     Tflag = 1 << 2
)
```

#### type Type

```go
type Type struct {
	Size       uintptr
	Ptrdata    uintptr // size of memory prefix holding all pointers
	Hash       uint32
	Tflag      Tflag
	Align      uint8
	Fieldalign uint8
	Kind       uint8
	Alg        *TypeAlg
	// gcdata stores the GC Type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	Gcdata    *byte
	Str       NameOff
	PtrToThis TypeOff
}
```

Needs to be in sync with
../cmd/compile/internal/ld/decodesym.go:/^func.commonsize,
../cmd/compile/internal/gc/reflect.go:/^func.dcommonType and
../reflect/Type.go:/^Type.rType.

#### func  PtrToType

```go
func PtrToType(ptr unsafe.Pointer) *Type
```

#### func  ResolveTypeOff

```go
func ResolveTypeOff(ptrInModule unsafe.Pointer, off TypeOff) *Type
```

#### func (*Type) Name

```go
func (t *Type) Name() string
```

#### func (*Type) NameOff

```go
func (t *Type) NameOff(off NameOff) Name
```

#### func (*Type) String

```go
func (t *Type) String() string
```

#### func (*Type) TextOff

```go
func (t *Type) TextOff(off textOff) unsafe.Pointer
```

#### func (*Type) TypeOff

```go
func (t *Type) TypeOff(off TypeOff) *Type
```

#### func (*Type) Uncommon

```go
func (t *Type) Uncommon() *UncommonType
```

#### type TypeAlg

```go
type TypeAlg struct {
	// function for hashing objects of this type
	// (ptr to object, seed) -> hash
	Hash func(unsafe.Pointer, uintptr) uintptr
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	Equal func(unsafe.Pointer, unsafe.Pointer) bool
}
```

TypeAlg is also copied/used in reflect/type.go. keep them in sync.

#### type TypeOff

```go
type TypeOff int32
```


#### type UncommonType

```go
type UncommonType struct {
	Pkgpath NameOff
	Mcount  uint16 // number of methods

	Moff uint32 // offset from this UncommonType to [mcount]method
}
```
