package errors

type fundamental struct {
	msg   string
	stack []uintptr
}

func (e *fundamental) Error() string { return e.msg }

func (e *fundamental) StackTrace() []uintptr { return e.stack }

func (e *fundamental) ErrorStack() string { return e.msg + "\n" + stackString(e.stack) }
