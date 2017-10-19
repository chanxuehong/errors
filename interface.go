package errors

type causer interface {
	Cause() error
}

type errorStacker interface {
	ErrorStack() string
}
