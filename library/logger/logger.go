package logger

import (
	"fmt"
	"myTool/mylog"
	"sync"
)

var (
	logger *mylog.DebugLogger
	once   sync.Once
)

func GetLogger(projectName string) *mylog.DebugLogger {
	once.Do(func() {
		logger = mylog.DefaultDebugLogger(projectName + "_log")
	})
	return logger
}

func Println(a ...interface{}) {
	fmt.Println(a)
}

func Printf(format string, a ...interface{}) {
	fmt.Printf(format, a)
}
