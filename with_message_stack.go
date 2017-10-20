package errors

import (
	"fmt"
	"io"
)

var (
	_ error         = (*withMessageStack)(nil)
	_ causer        = (*withMessageStack)(nil)
	_ errorStacker  = (*withMessageStack)(nil)
	_ fmt.Formatter = (*withMessageStack)(nil)
)

type withMessageStack struct {
	cause error
	msg   string
	stack []uintptr
}

func (e *withMessageStack) Error() string { return e.Cause().Error() + ": " + e.msg }

func (e *withMessageStack) Cause() error { return e.cause }

func (e *withMessageStack) ErrorStack() string {
	return ErrorStack(e.Cause()) + "\n" + e.msg + "\n" + stackString(e.stack)
}

func (e *withMessageStack) Format(s fmt.State, verb rune) {
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
