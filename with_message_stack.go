package errors

type withMessageStack struct {
	withMessage
	stack []uintptr
}

func (e *withMessageStack) StackTrace() []uintptr { return e.stack }

func (e *withMessageStack) ErrorStack() string {
	return e.withMessage.ErrorStack() + "\n" + stackString(e.stack)
}
