package main

import (
	"fmt"
	"log"
	"unsafe"
)

type runtimeTypeMimic struct {
	Size_       uintptr
	PtrBytes    uintptr // number of (prefix) bytes in the type that can contain pointers
	Hash        uint32  // hash of type; avoids computation in hash tables
	TFlag       uint8   // extra type information flags
	Align_      uint8   // alignment of variable with this type
	FieldAlign_ uint8   // alignment of struct field with this type
	Kind_       uint8   // enumeration for C
	Equal       func(unsafe.Pointer, unsafe.Pointer) bool
	GCData      *byte
	Str         int32 // string form
	PtrToThis   int32 // type for pointer to this type, may be zero
}

type eFaceMimic struct {
	Type *runtimeTypeMimic
	Data unsafe.Pointer
}

type TypeInfo struct {
	Hash  uint32
	TFlag uint8
	Kind  uint8
	Equal func(unsafe.Pointer, unsafe.Pointer) bool
	This  unsafe.Pointer
}

func GetTypeInfo(v any) TypeInfo {
	eface := *(*eFaceMimic)(unsafe.Pointer(&v))
	return TypeInfo{
		Hash:  eface.Type.Hash,
		TFlag: eface.Type.TFlag,
		Kind:  eface.Type.Kind_,
		Equal: eface.Type.Equal,
		This:  eface.Data,
	}
}

func main() {
	i := 42
	eface := GetTypeInfo(i)
	fmt.Printf("%+v\n", eface)

	f := 3.14
	eface2 := GetTypeInfo(f)
	fmt.Printf("%+v\n", eface2)

	// Compare the two values using the Equal function from the type info
	log.Println(eface.Equal(eface.This, eface.This))
	log.Println(eface2.Equal(eface2.This, eface2.This))
	log.Println(eface2.Equal(eface.This, eface2.This))
	log.Println(eface.Equal(eface2.This, eface.This))
}
