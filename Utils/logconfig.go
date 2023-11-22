package Utils

import (
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger // 包级别的日志记录器

func init() {
	logger = log.New()
	logger.SetFormatter(&log.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2006-01-02 15:04:05",
		FullTimestamp:             true,
		DisableLevelTruncation:    true,
	})
	logger.SetLevel(log.InfoLevel)
}

// GetLogger 获取日志记录器
func GetLogger() *log.Logger {
	return logger
}
