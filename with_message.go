package errors

import (
	"fmt"
	"io"
)

var (
	_ error         = (*withMessage)(nil)
	_ causer        = (*withMessage)(nil)
	_ errorStacker  = (*withMessage)(nil)
	_ fmt.Formatter = (*withMessage)(nil)
)

type withMessage struct {
	cause error
	msg   string
}

func (e *withMessage) Error() string { return e.Cause().Error() + ": " + e.msg }

func (e *withMessage) Cause() error { return e.cause }

func (e *withMessage) ErrorStack() string { return ErrorStack(e.Cause()) + "\n" + e.msg }

func (e *withMessage) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, e.ErrorStack())
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}
