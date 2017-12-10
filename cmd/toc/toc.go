// toc working with links and table of contents
package main

import (
	"flag"
	"github.com/gregoryv/find"
	"github.com/gregoryv/website"
	"golang.org/x/net/html"
	"log"
	"os"
	"path"
)

var (
	root string
	verbose bool
)

func init() {
	root, _ := os.Getwd()
	flag.StringVar(&root, "r", root, "Root directory")
	flag.BoolVar(&verbose, "verbose", verbose, "Log progress to stdout")
}

func main() {
	flag.Parse()
	var err error
	// Check links
	htmlFiles, err := find.By(find.NewShellPattern("*.html"), root)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range htmlFiles {
		log.Printf("Checking %s", file)
		fh, err := os.Open(file)
		if err != nil {
			log.Fatalf("%s", err)
		}
		defer fh.Close()
		doc, err := html.Parse(fh)
		if err != nil {
			log.Fatalf("%s", err)
		}
		website.CheckLink(path.Dir(file), doc)
	}
}
