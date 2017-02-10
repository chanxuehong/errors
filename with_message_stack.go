package errors

import "fmt"

func WithStack(err error) error {
	if err == nil {
		return nil
	}
	v, ok := err.(StackTracer)
	if !ok {
		return &withMessageStack{
			cause: err,
			msg:   "",
			stack: callers(),
		}
	}
	stack := v.StackTrace()
	if len(stack) == 0 {
		return &withMessageStack{
			cause: err,
			msg:   "",
			stack: callers(),
		}
	}
	return err // TODO 强制转换到这个包的类型?
}

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	v, ok := err.(StackTracer)
	if !ok {
		return &withMessageStack{
			cause: err,
			msg:   msg,
			stack: callers(),
		}
	}
	stack := v.StackTrace()
	if len(stack) == 0 {
		return &withMessageStack{
			cause: err,
			msg:   msg,
			stack: callers(),
		}
	}
	if msg == "" {
		return err // TODO 强制转换到这个包的类型?
	}
	return &withMessageStack{
		cause: err,
		msg:   msg,
		stack: stack,
	}
}

func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf(format, args...)
	v, ok := err.(StackTracer)
	if !ok {
		return &withMessageStack{
			cause: err,
			msg:   msg,
			stack: callers(),
		}
	}
	stack := v.StackTrace()
	if len(stack) == 0 {
		return &withMessageStack{
			cause: err,
			msg:   msg,
			stack: callers(),
		}
	}
	if msg == "" {
		return err // TODO 强制转换到这个包的类型?
	}
	return &withMessageStack{
		cause: err,
		msg:   msg,
		stack: stack,
	}
}

var _ error = (*withMessageStack)(nil)
var _ Causer = (*withMessageStack)(nil)
var _ StackTracer = (*withMessageStack)(nil)

type withMessageStack struct {
	cause error
	msg   string
	stack []uintptr
}

// implements error
func (w *withMessageStack) Error() string {
	if w.msg == "" {
		return w.cause.Error()
	}
	return w.msg + ": " + w.cause.Error()
}

// implements causer
func (w *withMessageStack) Cause() error {
	return w.cause
}

// implements causer
func (w *withMessageStack) IID_93FF6FA1EDC311E6B34F38C98633AC15() {}

// implements stackTracer
func (w *withMessageStack) StackTrace() []uintptr {
	return w.stack
}

// implements stackTracer
func (w *withMessageStack) IID_9BB74855EDC311E689C438C98633AC15() {}
