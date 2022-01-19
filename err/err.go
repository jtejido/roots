package err

import (
	"log"
	// "os"
)

type (
	ErrorHandlerType  = func(reason, file string, line, gsl_errno int)
	StreamHandlerType = func(label, file string, line int, reason string)
)

var (
	errorHandler ErrorHandlerType = nil
)

type RootsError interface {
	error
	Status() int
}

type rootsError struct {
	status  int
	message string
}

func New(status int, text string) *rootsError {
	return &rootsError{status, text}
}

func (err *rootsError) Status() int {
	return err.status
}

func (err *rootsError) Error() string {
	return err.message
}

func HandleError(reason, file string, line, gsl_errno int) {
	if errorHandler != nil {
		errorHandler(reason, file, line, gsl_errno)
		return
	}

	StreamPrintf("ERROR", file, line, reason)
	log.Printf("Default Roots error handler invoked.\n")
	panic(reason)
}

func SetErrorHandler(new_handler ErrorHandlerType) ErrorHandlerType {
	previous_handler := errorHandler
	errorHandler = new_handler
	return previous_handler
}

func SetErrorHandlerOff() ErrorHandlerType {
	previous_handler := errorHandler
	errorHandler = NoErrorHandler
	return previous_handler
}

func NoErrorHandler(reason, file string, line int, gsl_errno int) {
	/* do nothing */
	return
}
