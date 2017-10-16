package errors

import (
	"fmt"
	"strings"
)

// Wrap returns an error annotating err with a stack trace
// at the point Wrap was called, and the supplied messages.
// If err is a StackTracer, the result of Wrap will also have the same stack trace as err.
// If err is nil, Wrap returns nil.
func Wrap(err error, msg ...string) error {
	return wrap(err, strings.Join(msg, ": "), false)
}

// WrapWithStackAlways returns an error annotating err with a stack trace
// at the point WrapWithStackAlways is called, and the supplied message.
// If err is nil, WrapWithStackAlways returns nil.
func WrapWithStackAlways(err error, msg ...string) error {
	return wrap(err, strings.Join(msg, ": "), true)
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf was called, and the message fmt.Sprintf(format, args...).
// If err is a StackTracer, the result of Wrapf will also have the same stack trace as err.
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...interface{}) error {
	return wrap(err, fmt.Sprintf(format, args...), false)
}

// WrapfWithStackAlways returns an error annotating err with a stack trace
// at the point WrapfWithStackAlways is call, and the format specifier.
// If err is nil, WrapfWithStackAlways returns nil.
func WrapfWithStackAlways(err error, format string, args ...interface{}) error {
	return wrap(err, fmt.Sprintf(format, args...), true)
}

func wrap(err error, msg string, withStackAlways bool) error {
	if err == nil {
		return nil
	}
	if withStackAlways {
		return &withMessageStack{
			cause: err,
			msg:   msg,
			stack: callers(3),
		}
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
		return err // TODO convert to *withMessageStack if err type is not *withMessageStack or *fundamental?
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
