package errors

import "fmt"

func WithStack(err error) error {
	return wrap(err, "")
}

func Wrap(err error, msg string) error {
	return wrap(err, msg)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return wrap(err, fmt.Sprintf(format, args...))
}

func wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	v, ok := err.(StackTracer)
	if !ok {
		return &withMessageStack{
			cause: err,
			msg:   msg,
			stack: callers(3),
		}
	}
	stack := v.StackTrace()
	if len(stack) == 0 {
		return &withMessageStack{
			cause: err,
			msg:   msg,
			stack: callers(3),
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
var _ errorStacker = (*withMessageStack)(nil)

type withMessageStack struct {
	cause       error
	msg         string
	stack       []uintptr
	stackString string
}

// implements error
func (w *withMessageStack) Error() string {
	if w.msg == "" {
		return w.cause.Error()
	}
	return w.msg + ": " + w.cause.Error()
}

// implements Causer
func (w *withMessageStack) Cause() error {
	return w.cause
}

// implements Causer
func (w *withMessageStack) IID_93FF6FA1EDC311E6B34F38C98633AC15() {}

// implements StackTracer
func (w *withMessageStack) StackTrace() []uintptr {
	return w.stack
}

// implements StackTracer
func (w *withMessageStack) IID_9BB74855EDC311E689C438C98633AC15() {}

// implements errorStacker
func (w *withMessageStack) errorStack() string {
	if w.stackString == "" {
		w.stackString = stackString(w.stack)
	}
	return w.Error() + "\n" + w.stackString
}
