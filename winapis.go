package winapi

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

const INVALID_HANDLE_VALUE = ^windows.Handle(0)

var (
	kernel32                     = windows.NewLazySystemDLL("kernel32.dll")
	procVirtualAllocEx           = kernel32.NewProc("VirtualAllocEx")
	procHeapAlloc                = kernel32.NewProc("HeapAlloc")
	procCreateThread             = kernel32.NewProc("CreateThread")
	procCreateRemoteThread       = kernel32.NewProc("CreateRemoteThread")
	procCreateToolhelp32Snapshot = kernel32.NewProc("CreateToolhelp32Snapshot")
)

func VirtualAllocEx(hFunction *windows.LazyProc, hProcess windows.Handle, lpAddress uintptr, dwSize uintptr, flAllocationType uint32, flProtect uint32) (uintptr, error) {
	r0, _, _ := procVirtualAllocEx.Call(uintptr(hProcess), lpAddress, dwSize, uintptr(flAllocationType), uintptr(flProtect))
	if r0 == 0 {
		return 0, windows.GetLastError()
	}
	return uintptr(r0), nil
}

func HeapAlloc(hFunction *windows.LazyProc, hHeap windows.Handle, dwFlags uint32, dwBytes uintptr) (uintptr, error) {
	r0, _, e1 := procHeapAlloc.Call(uintptr(hHeap), uintptr(dwFlags), dwBytes)
	if r0 == 0 {
		return 0, fmt.Errorf("HeapAlloc failed. Possible causes: STATUS_NO_MEMORY or STATUS_ACCESS_VIOLATION: %s ", e1)
	}
	return uintptr(r0), nil
}

func CreateThread(hFunction *windows.LazyProc, lpThreadAttributes uintptr, dwStackSize uint32, lpStartAddress uintptr, lpParameter uintptr, dwCreationFlags uint32, lpThreadId *uint32) (windows.Handle, error) {
	r0, _, _ := procCreateThread.Call(lpThreadAttributes, uintptr(dwStackSize), lpStartAddress, lpParameter, uintptr(dwCreationFlags), uintptr(unsafe.Pointer(lpThreadId)))
	if r0 == 0 {
		return 0, fmt.Errorf("CreateThread failed: %d", windows.GetLastError())
	}
	return windows.Handle(r0), nil
}

func CreateRemoteThread(hFunction *windows.LazyProc, hProcess windows.Handle, lpThreadAttributes uintptr, dwStackSize uint32, lpStartAddress uintptr, lpParameter uintptr, dwCreationFlags uint32, lpThreadId *uint32) (windows.Handle, error) {
	r0, _, _ := procCreateRemoteThread.Call(uintptr(hProcess), lpThreadAttributes, uintptr(dwStackSize), lpStartAddress, lpParameter, uintptr(dwCreationFlags), uintptr(unsafe.Pointer(lpThreadId)))
	if r0 == 0 {
		return 0, fmt.Errorf("CreateRemoteThread failed: %d", windows.GetLastError())
	}
	return windows.Handle(r0), nil
}

func CreateToolhelp32Snapshot(hFunction *windows.LazyProc, dwFlags uint32, th32ProcessID uint32) (windows.Handle, error) {
	r0, _, _ := procCreateToolhelp32Snapshot.Call(uintptr(dwFlags), uintptr(th32ProcessID))
	if windows.Handle(r0) == INVALID_HANDLE_VALUE {
		return 0, fmt.Errorf("CreateToolhelp32Snapshot failed: %d", windows.GetLastError())
	}
	return windows.Handle(r0), nil
}
