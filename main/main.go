package main

import (
	"flag"
	"fmt"
	"github.com/hanghang/md2min"
	"io/ioutil"
	"os"
	"strings"
)

var mode string
var modeList []string
var output string

func pFail() {
	if err := recover(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: md2min [-nav=h2] name.md\n  name.md: markdown file name\n")
	flag.PrintDefaults()
}

func init() {
	modeList = []string{"none", "h1", "h2", "h3", "h4", "h5", "h6"}
	flag.StringVar(&mode, "nav", "none", `navigate level ["none", "h1", "h2", "h3", "h4", "h5", "h6"]`)
	flag.StringVar(&output, "output", "file", `output ["file", "stdout"]`)
}

func main() {
	defer pFail()

	flag.Parse()

	i := 0
	for ; i < len(modeList); i++ {
		if modeList[i] == mode {
			break
		}
	}

	if flag.NArg() < 1 || i >= len(modeList) {
		usage()
		return
	}

	name := flag.Args()[0]
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	md := md2min.New(mode)

	if output == "stdout" {
		err = md.Parse(input, os.Stdout)
	} else {
		newname := strings.TrimRight(name, ".md") + ".html"
		outfile, err := os.Create(newname)
		if err != nil {
			panic(err)
		}
		defer outfile.Close()
		err = md.Parse(input, outfile)
	}
	if err != nil {
		panic(err)
	}
}
