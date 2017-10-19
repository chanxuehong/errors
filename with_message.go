package errors

type withMessage struct {
	cause error
	msg   string
}

func (e *withMessage) Error() string { return e.msg + ": " + e.cause.Error() }

func (e *withMessage) Cause() error { return e.cause }

func (e *withMessage) StackTrace() []uintptr {
	if cause, ok := e.Cause().(StackTracer); ok {
		return cause.StackTrace()
	}
	return nil
}

func (e *withMessage) ErrorStack() string { return String(e.cause) + "\n" + e.msg }
