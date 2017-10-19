package errors

import (
	"fmt"
	"io"
)

var (
	_ error         = (*fundamental)(nil)
	_ errorStacker  = (*fundamental)(nil)
	_ fmt.Formatter = (*fundamental)(nil)
)

type fundamental struct {
	msg   string
	stack []uintptr
}

func (e *fundamental) Error() string { return e.msg }

func (e *fundamental) ErrorStack() string { return e.Error() + "\n" + stackString(e.stack) }

func (e *fundamental) Format(s fmt.State, verb rune) {
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
