package main

import (
	"github.com/jay13jay/fileMgr/argtools"
)

func main() {
	handler, _ := argtools.ReadArgs()
	handler.ProcessArgs()

	// fmt.Println("Dirs: ", handler.Dirs)
}
