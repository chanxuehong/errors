package errors

type Causer interface {
	Cause() error
}

type StackTracer interface {
	StackTrace() (stack []uintptr, ok bool)
}

type errorStacker interface {
	ErrorStack() string
}
