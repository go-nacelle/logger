package log

import (
	"fmt"
	"runtime"
	"strings"
)

func addCaller(fields LogFields, depth int) LogFields {
	if fields == nil {
		fields = LogFields{}
	}

	if _, ok := fields["caller"]; !ok {
		fields["caller"] = getCaller(depth)
	}

	return fields
}

func getCaller(depth int) string {
	for i := 3 + depth; ; i++ {
		_, file, line, _ := runtime.Caller(i)
		if file == "<autogenerated>" {
			continue
		}

		return fmt.Sprintf("%s:%d", trimPath(file), line)
	}
}

func trimPath(path string) string {
	// For details, see http://goo.gl/FL2U8s.

	if idx := strings.LastIndexByte(path, '/'); idx >= 0 {
		if idx := strings.LastIndexByte(path[:idx], '/'); idx >= 0 {
			// Keep everything after the penultimate separator.
			return path[idx+1:]
		}
	}

	return path
}
