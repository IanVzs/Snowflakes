package logging

import (
	"fmt"

	"github.com/IanVzs/Snowflakes/pkgs/setting"
)

// 运行路径+日志路径(并不是很好控制,可能以后还是会改成只logSavePath)
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

// 文件名.log
func getLogFileName() string {
	return fmt.Sprintf("%s.%s", setting.AppSetting.LogSaveName, setting.AppSetting.LogFileExt)
}

func getLogLevel() string {
	return fmt.Sprintf("%s", setting.ServerSetting.RunMode)
}
