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

func wrap(err error, msg string, withStackAlways bool) error {
	if err == nil {
		return nil
	}
	if withStackAlways {
		if msg != "" {
			err = &withMessage{
				cause: err,
				msg:   msg,
			}
		}
		return &withStack{
			cause: err,
			stack: callers(3),
		}
	}
	if v, ok := err.(StackTracer); ok {
		if stack, ok := v.StackTrace(); ok && len(stack) > 0 {
			if msg == "" {
				return err
			}
			return &withMessage{
				cause: err,
				msg:   msg,
			}
		}
	}
	if msg != "" {
		err = &withMessage{
			cause: err,
			msg:   msg,
		}
	}
	return &withStack{
		cause: err,
		stack: callers(3),
	}
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the Causer interface.
//
// If the error does not implement Causer interface, the original error will
// be returned.
// If the error is nil, nil will be returned without further investigation.
func Cause(err error) error {
	var (
		causer Causer
		ok     bool
	)
	for err != nil {
		causer, ok = err.(Causer)
		if !ok {
			break
		}
		err = causer.Cause()
	}
	return err
}

// String returns the error message of err.
// If err does not implement StackTracer interface, String returns err.Error(),
// else it returns a string that contains both the error message and the callstack.
// If err is nil, String returns "".
func String(err error) string {
	if err == nil {
		return ""
	}
	if v, ok := err.(errorStacker); ok {
		return v.ErrorStack()
	}
	v, ok := err.(StackTracer)
	if !ok {
		return err.Error()
	}
	stack, _ := v.StackTrace()
	if len(stack) == 0 {
		return err.Error()
	}
	return err.Error() + "\n" + stackString(stack)
}
