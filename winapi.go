package main

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

func VirtualAllocEx(hProcess windows.Handle, lpAddress uintptr, dwSize uintptr, flAllocationType uint32, flProtect uint32) (uintptr, error) {
	r0, _, e1 := windows.NewLazyDLL("kernel32.dll").NewProc("VirtualAllocEx").Call(uintptr(hProcess), lpAddress, dwSize, uintptr(flAllocationType), uintptr(flProtect))
	if r0 == 0 {
		return 0, e1
	}
	return uintptr(r0), nil
}

func HeapAlloc(hHeap windows.Handle, dwFlags uint32, dwBytes uintptr) (uintptr, error) {
	r0, _, e1 := windows.NewLazyDLL("kernel32.dll").NewProc("HeapAlloc").Call(uintptr(hHeap), uintptr(dwFlags), dwBytes)
	if r0 == 0 {
		return 0, e1
	}
	return uintptr(r0), nil
}

func CreateThread(lpThreadAttributes uintptr, dwStackSize uint32, lpStartAddress uintptr, lpParameter uintptr, dwCreationFlags uint32, lpThreadId *uint32) (windows.Handle, error) {
	r0, _, e1 := windows.NewLazyDLL("kernel32.dll").NewProc("CreateThread").Call(lpThreadAttributes, uintptr(dwStackSize), lpStartAddress, lpParameter, uintptr(dwCreationFlags), uintptr(unsafe.Pointer(lpThreadId)))
	if r0 == 0 {
		return 0, e1
	}
	return windows.Handle(r0), nil
}

func CreateToolhelp32Snapshot(dwFlags uint32, th32ProcessID uint32) (windows.Handle, error) {
	r0, _, e1 := windows.NewLazyDLL("kernel32.dll").NewProc("CreateToolhelp32Snapshot").Call(uintptr(dwFlags), uintptr(th32ProcessID))
	if r0 == 0 {
		return 0, e1
	}
	return windows.Handle(r0), nil
}
