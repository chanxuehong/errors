package errors

import "fmt"

func WithMessage(err error, msg string) error {
	if err == nil {
		return nil
	}
	if msg == "" {
		return err
	}
	v, ok := err.(StackTracer)
	if !ok {
		return &withMessage{
			cause: err,
			msg:   msg,
		}
	}
	stack := v.StackTrace()
	if len(stack) == 0 {
		return &withMessage{
			cause: err,
			msg:   msg,
		}
	}
	return &withMessageStack{
		cause: err,
		msg:   msg,
		stack: stack,
	}
}

func WithMessagef(err error, format string, args ...interface{}) error {
	return WithMessage(err, fmt.Sprintf(format, args...))
}

var _ error = (*withMessage)(nil)
var _ Causer = (*withMessage)(nil)

type withMessage struct {
	cause error
	msg   string
}

// implements error
func (w *withMessage) Error() string {
	return w.msg + ": " + w.cause.Error()
}

// implements Causer
func (w *withMessage) Cause() error {
	return w.cause
}

// implements Causer
func (w *withMessage) IID_93FF6FA1EDC311E6B34F38C98633AC15() {}
