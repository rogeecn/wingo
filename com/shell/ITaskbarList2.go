package shell

import (
	"github.com/rogeecn/wingo/util"
	"syscall"
	"unsafe"

	"github.com/rogeecn/wingo/win"
	"github.com/rogeecn/wingo/com/shell/shellvt"
	"github.com/rogeecn/wingo/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist2
type ITaskbarList2 struct{ ITaskbarList }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ITaskbarList2.Release().
//
// Example:
//
//  taskbl2 := shell.NewITaskbarList2(
//      win.CoCreateInstance(
//          shellco.CLSID_TaskbarList, nil,
//          co.CLSCTX_INPROC_SERVER,
//          shellco.IID_ITaskbarList2),
//  )
//  defer taskbl2.Release()
func NewITaskbarList2(base win.IUnknown) ITaskbarList2 {
	return ITaskbarList2{ITaskbarList: NewITaskbarList(base)}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist2-markfullscreenwindow
func (me *ITaskbarList2) MarkFullscreenWindow(hwnd win.HWND, fullScreen bool) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList2)(unsafe.Pointer(*me.Ptr())).MarkFullscreenWindow, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwnd), util.BoolToUintptr(fullScreen))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
