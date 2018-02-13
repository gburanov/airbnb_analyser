package main

import (
	"fmt"
	"runtime"
)

func smartError(err error) error {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Errorf("[%s][%d] : %s", file, line, err.Error())
}
