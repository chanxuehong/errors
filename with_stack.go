package errors

type withStack struct {
	cause error
	stack []uintptr
}

func (e *withStack) Error() string { return e.cause.Error() }

func (e *withStack) Cause() error { return e.cause }

func (e *withStack) StackTrace() []uintptr { return e.stack }

func (e *withStack) ErrorStack() string { return String(e.cause) + "\n" + stackString(e.stack) }
