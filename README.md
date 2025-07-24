# qslog

For english: [README.en.md](README.en.md)

轻量级 Go 语言日志库，基于标准库 `log` 封装，支持日志分级输出（DEBUG / INFO / WARN / ERROR）。

## 特性

- 日志等级支持：DEBUG、INFO、WARN、ERROR
- 纯文本格式，输出时间、等级、文件名和行号
- 简单易用，无第三方依赖
- 可配置全局日志等级

---

## 安装

```bash
go get github.com/li-qs/qslog
```

## 使用

```go
package main

import (
	"github.com/li-qs/qslog"
)

func main() {
	qslog.SetLevel(qslog.DebugLevel)

	qslog.Debug("...")
	qslog.Info("...")
	qslog.Warn("...")
	qslog.Error("...")

	qslog.Infof("用户 %s 已登录", "admin") // 打印：[2025-01-01 01:01:01] [INFO] 用户 admin 已登录
	qslog.Errorf("错误: %d", 500) // 打印：[2025-01-01 01:01:01] [ERROR] 错误: 500
}
```

## API

| 函数                   | 说明                |
| ---------------------- | ------------------- |
| `SetLevel(Level)`      | 设置日志级别        |
| `Debug(args...)`       | 打印 Debug 级别日志 |
| `Info(args...)`        | 打印 Info 级别日志  |
| `Warn(args...)`        | 打印 Warn 级别日志  |
| `Error(args...)`       | 打印 Error 级别日志 |
| `Debugf(fmt, args...)` | 格式化 Debug 日志   |
| `Infof(fmt, args...)`  | 格式化 Info 日志    |
| `Warnf(fmt, args...)`  | 格式化 Warn 日志    |
| `Errorf(fmt, args...)` | 格式化 Error 日志   |

## License

MIT License © 2025 li-qs
