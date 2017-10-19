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
	withMessage
	stack []uintptr
}

func (e *withMessageStack) ErrorStack() string {
	return e.withMessage.ErrorStack() + "\n" + stackString(e.stack)
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
