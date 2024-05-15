# go-winapi
This module contains functions to use Windows APIs that are missing in the Windows package from the standard library.

I left to the user the power to choose when to load which DLL and get the addresses to functions. You may need them in your code, who knows.
For this reason, this module requires you to pass to the functions windows.LazyProc structuctures, initialized with the addresses of the functions exported by the DLLs.

Load the DLL and get the function address:

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
###Usage example:

```
var kernel32 = windows.NewLazySystemDLL("kernel32.dll")
var procVirtualAllocEx = kernel32.NewProc("VirtualAllocEx")
)
func main(){
memAddress, err :=winapi.VirtualAllocEx(procVirtualAllocEx, 0, 0, 0, 0, 0,)
}
```
