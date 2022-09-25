package main

import (
	"fmt"
	"github.com/juju/loggo"
	"github.com/juju/loggo/loggocolor"
	"os"
)

func main() {

	const DefaultWriterName = "default"
	//panic("eee")
	fmt.Println("hello")
	writer, err := loggo.RemoveWriter("default")
	loggo.RegisterWriter("default", writer)
	loggo.ReplaceDefaultWriter(loggocolor.NewWriter(os.Stderr))
	logger := loggo.GetLogger("foo.bar")
	fmt.Println(logger)
	fmt.Println(loggo.LoggerInfo())
	fmt.Println(err)
	logger.Debugf("sds")
}
