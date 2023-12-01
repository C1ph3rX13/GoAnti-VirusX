package AntiSandBox

import (
	"GoAnti-VirusX/API"
	"time"
)

func BootTime() (int, error) {
	api, _ := API.SetWindowsAPI()
	// 使用NewProc函数创建GetTickCount的进程
	GetTickCount := api.Kernel32.NewProc("GetTickCount")
	// 调用GetTickCount函数获取系统启动时间（以毫秒为单位）
	startTime, _, _ := GetTickCount.Call()
	// 如果返回值为0，表示获取失败，返回0和nil错误
	if startTime == 0 {
		return 0, nil
	}

	// 将毫秒转换为time.Duration类型
	checkTime := time.Duration(startTime * 1000 * 1000)
	// 定义一个时间阈值为30分钟
	setTime := 30 * time.Minute
	// 如果系统启动时间小于30分钟, 则返回0和nil错误, 否则返回1和nil错误
	if checkTime < setTime {
		return 0, nil
	} else {
		return 1, nil
	}
}
