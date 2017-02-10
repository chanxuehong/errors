package errors

import (
	"bytes"
	"runtime"
	"strconv"
)

type Causer interface {
	IID_93FF6FA1EDC311E6B34F38C98633AC15()

	error
	Cause() error
}

func Cause(err error) error {
	var (
		cause Causer
		ok    bool
	)
	for err != nil {
		cause, ok = err.(Causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

type StackTracer interface {
	IID_9BB74855EDC311E689C438C98633AC15()

	error
	StackTrace() []uintptr
}

func String(err error) string {
	// TODO 优化这里
	if err == nil {
		return ""
	}
	v, ok := err.(StackTracer)
	if !ok {
		return err.Error()
	}
	stack := v.StackTrace()
	if len(stack) == 0 {
		return err.Error()
	}
	frames := runtime.CallersFrames(stack)

	var buf bytes.Buffer
	buf.WriteString(err.Error())

	var (
		frame runtime.Frame
		more  bool
	)
	for {
		frame, more = frames.Next()
		if frame.Function != "" {
			buf.WriteByte('\n')
			buf.WriteString(frame.Function)
			buf.WriteString("\n\t")
			buf.WriteString(frame.File)
			buf.WriteByte(':')
			buf.WriteString(strconv.Itoa(frame.Line))
		}
		if !more {
			break
		}
	}

	return buf.String()
}
