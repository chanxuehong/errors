package errors

type Causer interface {
	IID_93FF6FA1EDC311E6B34F38C98633AC15()

	error
	Cause() error
}

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

type StackTracer interface {
	IID_9BB74855EDC311E689C438C98633AC15()

	error
	StackTrace() []uintptr
}
