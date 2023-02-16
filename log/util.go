/**
 * Copyright 2022-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package log

import (
	"fmt"
	"runtime"
	"strings"
)

// caller follows path to remove logrus and local log package rows
func caller() func(*runtime.Frame) (function string, file string) {
	// Range to look through the frame and remove the log package rows
	frameStart := 2
	frameEnd := 10
	return func(f *runtime.Frame) (function string, file string) {
		pc, file, line, ok := runtime.Caller(2)
		funcName := runtime.FuncForPC(pc)
		for i := frameStart; i < frameEnd; i++ {
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
