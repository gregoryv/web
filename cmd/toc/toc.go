// toc working with links and table of contents
package main

import (
	"flag"
	"github.com/gregoryv/web/site"
	"os"
	"log"
)

var (
	root    string
	verbose bool
)

func init() {
	root, _ := os.Getwd()
	flag.StringVar(&root, "r", root, "Root directory")
	flag.BoolVar(&verbose, "verbose", verbose, "Log progress to stdout")
}

func main() {
	flag.Parse()
	done := make(chan bool)
	broken := make(chan site.BrokenLink)
	go func() {
		for lnk := range broken {
			log.Printf("%s", lnk.String())
		}
		done <- true
	}()
	site.CheckLinks(root, broken)
	<- done
}
