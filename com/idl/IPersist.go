package idl

import (
	"syscall"
	"unsafe"

	"github.com/rogeecn/wingo/win"
	"github.com/rogeecn/wingo/com/idl/idlvt"
	"github.com/rogeecn/wingo/errco"
)

// IPersist COM interface.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ipersist
type IPersist struct{ win.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IPersist.Release().
func NewIPersist(base win.IUnknown) IPersist {
	return IPersist{IUnknown: base}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ipersist-getclassid
func (me *IPersist) GetClassID() *win.GUID {
	clsid := &win.GUID{}
	ret, _, _ := syscall.Syscall(
		(*idlvt.IPersist)(unsafe.Pointer(*me.Ptr())).GetClassID, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(clsid)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return clsid
	} else {
		panic(hr)
	}
}
