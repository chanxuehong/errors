package errors

func Cause(err error) error {
	var (
		cause Causer
		ok    bool
	)
	for err != nil {
		cause, ok = err.(Causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

type Causer interface {
	IID_93FF6FA1EDC311E6B34F38C98633AC15()

	error
	Cause() error
}

func String(err error) string {
	if err == nil {
		return ""
	}
	v, ok := err.(StackTracer)
	if !ok {
		return err.Error()
	}
	stack := v.StackTrace()
	if len(stack) == 0 {
		return err.Error()
	}
	if v, ok := err.(errorStacker); ok {
		return v.errorStack()
	}
	return err.Error() + "\n" + stackString(stack)
}

type StackTracer interface {
	IID_9BB74855EDC311E689C438C98633AC15()

	error
	StackTrace() []uintptr
}

type errorStacker interface {
	errorStack() string
}
