package AntiSandBox

import (
	"runtime"
)

func GetLogicalCPUCores() (int, error) {
	CountOfCPU := runtime.NumCPU()
	if CountOfCPU < 4 {
		return 0, nil
	}
	return 1, nil

}
