package main

import (
	"fmt"
	"os"
	"io/ioutil"

	"github.com/tkennes/jengo/cmd"
	"github.com/tkennes/jengo/src"
)

func main() {
	jengo_src.InitLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
