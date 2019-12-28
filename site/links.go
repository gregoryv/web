package site

import (
	"fmt"
	"github.com/gregoryv/find"
	"golang.org/x/net/html"
	"net/url"
	"os"
	"path"
)

type BrokenLink struct {
	File string
	Ref  string
	Err  error
}

func (l *BrokenLink) String() string {
	return fmt.Sprintf("%s -> %s: %s", l.File, l.Ref, l.Err)
}

func CheckLinks(root string, broken chan BrokenLink) {
	htmlFiles, _ := find.By(find.NewShellPattern("*.html"), root)
	for e := htmlFiles.Front(); e != nil; e = e.Next() {
		file, _ := e.Value.(string)
		fh, _ := os.Open(file)
		defer fh.Close()
		doc, _ := html.Parse(fh)
		CheckLink(file, path.Dir(file), doc, broken)
	}
	close(broken)
}

func CheckLink(file, rel string, n *html.Node, broken chan BrokenLink) {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" {
				url, _ := url.Parse(a.Val)
				switch url.Scheme {
				case "file", "":
					if _, err := os.Stat(path.Join(rel, url.Path)); err != nil {
						broken <- BrokenLink{file, url.Path, err}
					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		CheckLink(file, rel, c, broken)
	}
}
