// toc working with links and table of contents
package main

import (
	"flag"
	"github.com/gregoryv/web/site"
	"log"
	"os"
)

var (
	root    string
	verbose bool
	help    bool
)

func init() {
	root = "."
	flag.BoolVar(&help, "h", help, "Show this help and exit")
	flag.StringVar(&root, "r", root, "Root directory")
	flag.BoolVar(&verbose, "verbose", verbose, "Log progress to stdout")
}

func main() {
	flag.Parse()
	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}
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
