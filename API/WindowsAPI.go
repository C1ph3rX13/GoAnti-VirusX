package API

import (
	"syscall"
)

type WindowsAPI struct {
	Kernel32 *syscall.LazyDLL
	Ntdll    *syscall.LazyDLL
	Rpcrt4   *syscall.LazyDLL
}

func SetWindowsAPI() (*WindowsAPI, error) {
	api := &WindowsAPI{
		Kernel32: syscall.NewLazyDLL("kernel32.dll"),
		Ntdll:    syscall.NewLazyDLL("ntdll.dll"),
		Rpcrt4:   syscall.NewLazyDLL("Rpcrt4.dll"),
	}
	return api, nil
}

func (api *WindowsAPI) VirtualAlloc() *syscall.LazyProc {
	return api.Kernel32.NewProc("VirtualAlloc")
}

func (api *WindowsAPI) VirtualProtect() *syscall.LazyProc {
	return api.Kernel32.NewProc("VirtualProtect")
}

func (api *WindowsAPI) RtlCopyMemory() *syscall.LazyProc {
	return api.Ntdll.NewProc("RtlCopyMemory")
}

func (api *WindowsAPI) RtlCopyBytes() *syscall.LazyProc {
	return api.Ntdll.NewProc("RtlCopyBytes")
}

func (api *WindowsAPI) ConvertThreadToFiber() *syscall.LazyProc {
	return api.Kernel32.NewProc("ConvertThreadToFiber")
}

func (api *WindowsAPI) CreateFiber() *syscall.LazyProc {
	return api.Kernel32.NewProc("CreateFiber")
}

func (api *WindowsAPI) SwitchToFiber() *syscall.LazyProc {
	return api.Kernel32.NewProc("SwitchToFiber")
}

func (api *WindowsAPI) GetCurrentThread() *syscall.LazyProc {
	return api.Kernel32.NewProc("GetCurrentThread")
}

func (api *WindowsAPI) NtQueueApcThreadEx() *syscall.LazyProc {
	return api.Ntdll.NewProc("NtQueueApcThreadEx")
}

func (api *WindowsAPI) EtwpCreateEtwThread() *syscall.LazyProc {
	return api.Ntdll.NewProc("EtwpCreateEtwThread")
}

func (api *WindowsAPI) WaitForSingleObject() *syscall.LazyProc {
	return api.Kernel32.NewProc("WaitForSingleObject")
}

func (api *WindowsAPI) CreateThread() *syscall.LazyProc {
	return api.Kernel32.NewProc("CreateThread")
}

func (api *WindowsAPI) OpenProcess() *syscall.LazyProc {
	return api.Kernel32.NewProc("OpenProcess")
}

func (api *WindowsAPI) VirtualAllocEx() *syscall.LazyProc {
	return api.Kernel32.NewProc("VirtualAllocEx")
}

func (api *WindowsAPI) VirtualProtectEx() *syscall.LazyProc {
	return api.Kernel32.NewProc("VirtualProtectEx")
}

func (api *WindowsAPI) WriteProcessMemory() *syscall.LazyProc {
	return api.Kernel32.NewProc("WriteProcessMemory")
}

func (api *WindowsAPI) CreateRemoteThreadEx() *syscall.LazyProc {
	return api.Kernel32.NewProc("CreateRemoteThreadEx")
}

func (api *WindowsAPI) CloseHandle() *syscall.LazyProc {
	return api.Kernel32.NewProc("CloseHandle")
}

func (api *WindowsAPI) HeapCreate() *syscall.LazyProc {
	return api.Kernel32.NewProc("HeapCreate")
}

func (api *WindowsAPI) HeapAlloc() *syscall.LazyProc {
	return api.Kernel32.NewProc("HeapAlloc")
}

func (api *WindowsAPI) EnumSystemLocalesA() *syscall.LazyProc {
	return api.Kernel32.NewProc("EnumSystemLocalesA")
}

func (api *WindowsAPI) UuidFromStringA() *syscall.LazyProc {
	return api.Rpcrt4.NewProc("UuidFromStringA")
}
