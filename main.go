package main

import (
	"os"

	"go.uber.org/zap"
)

func main() {
	fname := os.Args[1]

	lgr, _ := zap.NewDevelopment()
	log := lgr.Sugar()

	log.Debugw("starting", "filename", fname)
}
