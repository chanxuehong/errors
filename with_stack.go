package errors

type withStack struct {
	cause error
	stack []uintptr
}

func (w *withStack) Error() string { return w.cause.Error() }

func (w *withStack) Cause() error { return w.cause }

func (w *withStack) StackTrace() ([]uintptr, bool) { return w.stack, true }

func (w *withStack) ErrorStack() string {
	return String(w.cause) + "\n" + stackString(w.stack)
}
