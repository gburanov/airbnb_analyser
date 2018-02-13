package main

import (
	"fmt"
	"runtime"
)

type smartError struct {
	subError error
	errType  string
	file     string
	line     int
}

func newSmartError(err error, errType string) *smartError {
	_, file, line, _ := runtime.Caller(1)
	return &smartError{
		subError: err,
		errType:  errType,
		file:     file,
		line:     line,
	}
}

func (s *smartError) getType() string {
	return ""
}

func (s *smartError) Error() string {
	return fmt.Sprintf("[%s][%d] : %s", s.file, s.line, s.subError.Error())
}
