package main

import "fmt"

type writer interface {
	Write(str string) error
}

type outputWriter struct {
}

func (w *outputWriter) Write(str string) error {
	fmt.Println(str)
	return nil
}
