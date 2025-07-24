package main

import (
	"github.com/li-qs/qslog"
)

func main() {
	// 设置日志等级（默认是 Info，可选 DebugLevel、WarnLevel、ErrorLevel）
	qslog.SetLevel(qslog.DebugLevel)

	// 示例输出
	qslog.Debug("This is a debug message")
	qslog.Info("Hello from qslog!")
	qslog.Warn("This is a warning")
	qslog.Error("This is an error")

	// 格式化日志
	qslog.Infof("Current user: %s", "admin")
	qslog.Errorf("Status code: %d", 500)
}
