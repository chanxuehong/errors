package errors

import (
	"fmt"
	"io"
)

var (
	_ error         = (*withStack)(nil)
	_ causer        = (*withStack)(nil)
	_ errorStacker  = (*withStack)(nil)
	_ fmt.Formatter = (*withStack)(nil)
)

type withStack struct {
	cause error
	stack []uintptr
}

func (e *withStack) Error() string { return e.Cause().Error() }

func (e *withStack) Cause() error { return e.cause }

func (e *withStack) ErrorStack() string { return ErrorStack(e.Cause()) + "\n" + stackString(e.stack) }

func (e *withStack) Format(s fmt.State, verb rune) {
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
