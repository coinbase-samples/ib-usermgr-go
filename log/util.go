package log

import (
	"fmt"
	"runtime"
	"strings"
)

// caller follows path to remove logrus and local log package rows
func caller() func(*runtime.Frame) (function string, file string) {
	return func(f *runtime.Frame) (function string, file string) {
		pc, file, line, ok := runtime.Caller(2)
		funcName := runtime.FuncForPC(pc)
		for i := 2; i < 10; i++ {
			if ok && (strings.Contains(file, "entry.go") ||
				strings.Contains(file, "log.go") ||
				strings.Contains(file, "json_formatter.go") ||
				strings.Contains(file, "logger.go")) {
				pc, file, line, ok = runtime.Caller(i)
				funcName = runtime.FuncForPC(pc)
			} else {
				break
			}
		}

		if !ok {
			return f.Function, fmt.Sprintf("%s:%d", f.File, f.Line)
		} else {
			return funcName.Name(), fmt.Sprintf("%s:%d", file, line)
		}
	}
}
