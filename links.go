package website

import (
	"log"
	"golang.org/x/net/html"
	"strings"
	"os"
	"path"
)

func CheckLink(rel string, n *html.Node) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			for _, a := range n.Attr {
				if a.Key == "href" {
					linkOk(rel, a.Val)
				}
			}
		case "img", "link":
			for _, a := range n.Attr {
				if a.Key == "src" {
					linkOk(rel, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		CheckLink(rel, c)
	}
}

func linkOk(rel, file string) {
	if file != "" && strings.Index(file, "http") != 0 && strings.Index(file, "#") != 0 {
		if _, err := os.Stat(path.Join(rel, file)); err != nil {
			log.Print(err)
		}
	}
}
