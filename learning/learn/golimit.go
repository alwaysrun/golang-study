package learn

import (
	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/gatefs"
)


func LimitTest()  {
	fs :=vfs.OS("/path")
	gate := gatefs.New(fs, make(chan bool, 5))

	gate.Open("first")

}
