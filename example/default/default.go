package main

import (
	"os"

	"github.com/ddo/go-dlog"
)

// if nil writer -> os.Stdout
func main() {
	logDefault := dlog.New(nil)

	logDefault("logDefault 1")
	logDefault("logDefault 2")

	logStdout := dlog.New(&dlog.Option{
		Name:   "logStdout",
		Writer: os.Stdout,
	})

	logStdout("logStdout 1")
	logStdout("logStdout 2")
}