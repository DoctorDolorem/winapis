# go-winapis
Module that contains some functions to use WinAPIs that are missing in the main Windows package

I left to the user the power to choose when to load which DLL and get the addresses to functions. For this reason,
the functions require a pointer to a windows.LazyProc structure, which is the associated function exported by the dll.

DLL's must be loaded beforehand by the user:
```
var (
	kernel32                     = windows.NewLazySystemDLL("kernel32.dll")
	procVirtualAllocEx           = kernel32.NewProc("VirtualAllocEx")
	procHeapAlloc                = kernel32.NewProc("HeapAlloc")
	procCreateThread             = kernel32.NewProc("CreateThread")
	procCreateRemoteThread       = kernel32.NewProc("CreateRemoteThread")
	procCreateToolhelp32Snapshot = kernel32.NewProc("CreateToolhelp32Snapshot")
)
```
