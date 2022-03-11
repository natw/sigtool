package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/natw/sigtool/codesign"
)

func main() {
	fname := os.Args[1]

	lgr, _ := zap.NewDevelopment()
	log := lgr.Sugar()

	log.Debugw("starting", "filename", fname)

	cs := codesign.NewCodesign(&codesign.Options{})
	results, err := cs.VerifyRecursive(fname)
	if err != nil {
		log.Fatalw("oops", "err", err)
	}

	fmt.Printf("%+v\n", results)
}
