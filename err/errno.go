package err

import (
	"runtime"
)

const (
	EDOM = iota
	ERANGE
	EINVAL
	EBADFUNC
	EMAXITER
)

//  call the error handler, and return the error
func Error(reason string, errno int) RootsError {
	_, file, line, _ := runtime.Caller(2)
	HandleError(reason, file, line, errno)
	return New(errno, reason)
}
