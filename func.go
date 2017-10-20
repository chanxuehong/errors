package errors

import (
	"fmt"
	"strings"
)

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(msg string) error {
	return &fundamental{
		msg:   msg,
		stack: callers(2),
	}
}

// Newf returns an error with the message fmt.Sprintf(format, args...).
// Newf also records the stack trace at the point it was called.
func Newf(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return &fundamental{
		msg:   msg,
		stack: callers(2),
	}
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap was called if err does not implement StackTracer, and the supplied messages.
// If err is a StackTracer, the result of Wrap will also have the same stack trace as err.
// If err is nil, Wrap returns nil.
func Wrap(err error, msg ...string) error {
	return wrap(err, strings.Join(msg, ": "), false)
}

// WrapWithCurrentStackAlways returns an error annotating err with a stack trace
// at the point WrapWithCurrentStackAlways is called, and the supplied message.
// If err is nil, WrapWithCurrentStackAlways returns nil.
func WrapWithCurrentStackAlways(err error, msg ...string) error {
	return wrap(err, strings.Join(msg, ": "), true)
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf was called if err does not implement StackTracer, and the message fmt.Sprintf(format, args...).
// If err is a StackTracer, the result of Wrapf will also have the same stack trace as err.
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...interface{}) error {
	return wrap(err, fmt.Sprintf(format, args...), false)
}

// WrapfWithCurrentStackAlways returns an error annotating err with a stack trace
// at the point WrapfWithCurrentStackAlways is call, and the format specifier.
// If err is nil, WrapfWithCurrentStackAlways returns nil.
func WrapfWithCurrentStackAlways(err error, format string, args ...interface{}) error {
	return wrap(err, fmt.Sprintf(format, args...), true)
}

func wrap(err error, msg string, withCurrentStackAlways bool) error {
	if err == nil {
		return nil
	}
	if !withCurrentStackAlways {
		if _, ok := err.(errorStacker); ok {
			if msg == "" {
				return err
			}
			return &withMessage{
				cause: err,
				msg:   msg,
			}
		}
	}
	if msg == "" {
		return &withStack{
			cause: err,
			stack: callers(3),
		}
	}
	return &withMessageStack{
		cause: err,
		msg:   msg,
		stack: callers(3),
	}
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the causer interface.
//
// If the error does not implement causer interface, the original error will be returned.
// If the error is nil, nil will be returned without further investigation.
func Cause(err error) error {
	if err == nil {
		return nil
	}
	for {
		causer, ok := err.(causer)
		if !ok {
			return err
		}
		cause := causer.Cause()
		if cause == nil {
			return err
		}
		err = cause
	}
}

// ErrorStack returns the error message of err.
// If err does not implement errorStacker interface, ErrorStack returns err.Error(),
// else it returns a string that contains both the error message and the callstack.
// If err is nil, ErrorStack returns "".
func ErrorStack(err error) string {
	if err == nil {
		return ""
	}
	if v, ok := err.(errorStacker); ok {
		return v.ErrorStack()
	}
	return err.Error()
}
