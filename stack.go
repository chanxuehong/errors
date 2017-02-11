package errors

import (
	"bytes"
	"runtime"
	"strconv"
)

func callers(skip int) []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip+1, pcs[:])
	return pcs[:n]
}

func stackString(stack []uintptr) string {
	if len(stack) == 0 {
		return ""
	}
	frames := runtime.CallersFrames(stack)

	var (
		buf   bytes.Buffer
		frame runtime.Frame
		more  bool
	)
	for {
		frame, more = frames.Next()
		if frame.Function == "" {
			frame.Function = "unknown_function"
		}
		if frame.File == "" {
			frame.File = "unknown_file"
		}
		if buf.Len() > 0 {
			buf.WriteByte('\n')
		}
		buf.WriteString(frame.Function)
		buf.WriteString("\n\t")
		buf.WriteString(frame.File)
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(frame.Line))

		if !more {
			break
		}
	}
	return buf.String()
}
