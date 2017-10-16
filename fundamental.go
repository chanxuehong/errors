package errors

type fundamental struct {
	msg   string
	stack []uintptr
}

func (f *fundamental) Error() string { return f.msg }

func (f *fundamental) StackTrace() ([]uintptr, bool) { return f.stack, true }

func (f *fundamental) ErrorStack() string { return f.msg + "\n" + stackString(f.stack) }
