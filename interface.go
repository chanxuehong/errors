package errors

type Causer interface {
	Cause() error
}

type StackTracer interface {
	StackTrace() (stack []uintptr)
}

type errorStacker interface {
	ErrorStack() string
}
