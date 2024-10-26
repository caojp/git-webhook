package logger

import (
	"log"
	"os"
	"path/filepath"
)

var Logger *log.Logger

// Init initializes the logger and creates the log directory if it does not exist.
func Init(logFilePath string) error {
	// 获取日志文件的目录
	logDir := filepath.Dir(logFilePath)

	// 检查目录是否存在，如果不存在则创建
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0755) // 创建目录及其父目录
		if err != nil {
			return err
		}
		log.Println("Log directory created:", logDir) // 输出创建的日志目录
	} else {
		log.Println("Log directory exists:", logDir) // 输出已存在的日志目录
	}

	// 打开日志文件
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// 初始化日志记录器
	Logger = log.New(logFile, "", log.LstdFlags|log.Lshortfile)
	return nil
}
