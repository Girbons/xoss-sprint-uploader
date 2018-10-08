package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Girbons/xoss-uploader/uploader"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	File    string `short:"f" long:"file" required:"true" description:"The file that should be uploaded"`
	Private bool   `short:"p" long:"private" description:"Upload the activity as Private"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	data, err := ioutil.ReadFile(options.File)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	uploader.UploadFitFile(string(data), options.Private)
}
