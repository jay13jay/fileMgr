package main

import (
	"github.com/jay13jay/fileMgr/argtools"
)


func main() {
	handler,_ := argtools.ReadArgs()
	handler.ProcessArgs()
}