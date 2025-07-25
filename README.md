# qslog

中文: [README.zh.md](README.zh.md)

A lightweight Go logging library built on the standard `log` package, supporting leveled logs (DEBUG / INFO / WARN / ERROR).

## Features

- Log levels: DEBUG, INFO, WARN, ERROR
- Plain text output with timestamp, level, file, and line number
- Simple usage with zero dependencies
- Global log level configuration

---

## Installation

```bash
go get github.com/li-qs/qslog
```

## Usage

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

	qslog.Infof("User %s has logged in", "admin")
	qslog.Errorf("Error code: %d", 500)
}
````

## API

| Function               | Description                   |
| ---------------------- | ----------------------------- |
| `SetLevel(Level)`      | Set the logging level         |
| `Debug(args...)`       | Log a message at Debug level  |
| `Info(args...)`        | Log a message at Info level   |
| `Warn(args...)`        | Log a message at Warn level   |
| `Error(args...)`       | Log a message at Error level  |
| `Debugf(fmt, args...)` | Log a formatted Debug message |
| `Infof(fmt, args...)`  | Log a formatted Info message  |
| `Warnf(fmt, args...)`  | Log a formatted Warn message  |
| `Errorf(fmt, args...)` | Log a formatted Error message |

## License

MIT License © 2025 li-qs
