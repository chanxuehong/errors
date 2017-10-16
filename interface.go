package errors

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the Causer interface.
//
// If the error does not implement Causer interface, the original error will
// be returned.
// If the error is nil, nil will be returned without further investigation.
func Cause(err error) error {
	var (
		causer Causer
		ok     bool
	)
	for err != nil {
		causer, ok = err.(Causer)
		if !ok {
			break
		}
		err = causer.Cause()
	}
	return err
}

type Causer interface {
	IID_93FF6FA1EDC311E6B34F38C98633AC15()

	Cause() error
}

// String returns the error message of err.
// If err does not implement StackTracer interface, String returns err.Error(),
// else it returns a string that contains both the error message and the callstack.
// If err is nil, String returns "".
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

	StackTrace() []uintptr
}

type errorStacker interface {
	errorStack() string
}
