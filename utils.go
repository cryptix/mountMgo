package main

import (
	"fmt"
	"os"
	"runtime"
)

func checkFatal(err error) {
	_, file, line, _ := runtime.Caller(1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal from <%s:%d>\nError:%s", file, line, err)
		os.Exit(1)
	}
}

func logErr(err error) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "Error from <%s:%d>\nError:%s\n", file, line, err)
}
