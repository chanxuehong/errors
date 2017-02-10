package errors

import "fmt"

func New(msg string) error {
	return &fundamental{
		msg:   msg,
		stack: callers(2),
	}
}

func Errorf(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return &fundamental{
		msg:   msg,
		stack: callers(2),
	}
}

var _ error = (*fundamental)(nil)
var _ StackTracer = (*fundamental)(nil)

type fundamental struct {
	msg   string
	stack []uintptr
}

// implements error
func (f *fundamental) Error() string {
	return f.msg
}

// implements stackTracer
func (f *fundamental) StackTrace() []uintptr {
	return f.stack
}

// implements stackTracer
func (f *fundamental) IID_9BB74855EDC311E689C438C98633AC15() {}
