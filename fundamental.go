package errors

import "fmt"

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(msg string) error {
	return &fundamental{
		msg:   msg,
		stack: callers(2),
	}
}

// Errorf returns an error with the message fmt.Sprintf(format, args...).
// Errorf also records the stack trace at the point it was called.
func Errorf(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return &fundamental{
		msg:   msg,
		stack: callers(2),
	}
}

var _ error = (*fundamental)(nil)
var _ StackTracer = (*fundamental)(nil)
var _ errorStacker = (*fundamental)(nil)

type fundamental struct {
	msg         string
	stack       []uintptr
	stackString string
}

// implements error
func (f *fundamental) Error() string {
	return f.msg
}

// implements StackTracer
func (f *fundamental) StackTrace() []uintptr {
	return f.stack
}

// implements StackTracer
func (f *fundamental) IID_9BB74855EDC311E689C438C98633AC15() {}

// implements errorStacker
func (f *fundamental) errorStack() string {
	if f.stackString == "" {
		f.stackString = stackString(f.stack)
	}
	return f.Error() + "\n" + f.stackString
}
