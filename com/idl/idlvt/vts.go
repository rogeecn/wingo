package idlvt

import (
	"github.com/rogeecn/wingo/win"
)

// IPersist virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ipersist
type IPersist struct {
	win.IUnknownVtbl
	GetClassID uintptr
}
