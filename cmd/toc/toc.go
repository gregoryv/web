// toc working with links and table of contents
package main

import (
	"flag"
	"fmt"
	"github.com/gregoryv/web/site"
	"log"
	"os"
)

var help = false

func init() {
	flag.BoolVar(&help, "h", help, "Show this help and exit")
}

func usage() {
	if help {
		fmt.Printf("Usage: %s [OPTIONS] [PATH]\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func parsePath() string {
	args := flag.Args()
	if len(args) == 1 {
		return args[0]
	}
	return "."
}

func main() {
	flag.Parse()
	usage()
	root := parsePath()

	done := make(chan bool)
	broken := make(chan site.BrokenLink)
	var isBroken bool
	go func() {
		for lnk := range broken {
			log.Printf("%s", lnk.String())
			isBroken = true
		}
		done <- true
	}()
	site.CheckLinks(root, broken)
	<-done
	if isBroken {
		os.Exit(1)
	}
}
