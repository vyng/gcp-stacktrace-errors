package stacktrace

import (
	"cloud.google.com/go/errorreporting"
	"runtime/debug"
)

// WrapError creates a stack trace and returns an errorreporting entry
// TODO: This probably should not be in this repo so that we can keep dependencies out. But it is convenient!
func WrapError(err error, skip int) errorreporting.Entry {
	return errorreporting.Entry{
		Error: err,
		Stack: Stacktrace(debug.Stack(), skip + 1),
	}
}