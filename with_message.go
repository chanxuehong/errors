package errors

type withMessage struct {
	cause error
	msg   string
}

func (w *withMessage) Error() string { return w.msg + ": " + w.cause.Error() }

func (w *withMessage) Cause() error { return w.cause }

func (w *withMessage) StackTrace() ([]uintptr, bool) {
	if cause, ok := w.Cause().(StackTracer); ok {
		return cause.StackTrace()
	}
	return nil, false
}

func (w *withMessage) ErrorStack() string {
	return String(w.cause) + "\n" + w.msg
}
