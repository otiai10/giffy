package main

import "fmt"

type logger struct {
	quiet bool
}

func (l logger) Printf(format string, v ...interface{}) {
	if quiet {
		return
	}
	fmt.Printf(format, v...)
}
