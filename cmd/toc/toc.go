// toc working with links and table of contents
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gregoryv/web"
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

	err := web.CheckLinks(root)
	if err != nil {
		log.Fatal(err)
	}
}
