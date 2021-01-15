package web

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/gregoryv/find"
	"golang.org/x/net/html"
)

// ILinkAll replaces key words same as LinkAll but is canse insensitive
func ILinkAll(root *Element, refs map[string]string) {
	linkAll(root, refs, false)
}

// LinkAll replaces key words found in the root and it's
// children with links defined in the map. The map should be TEXT -> HREF
func LinkAll(root *Element, refs map[string]string) {
	linkAll(root, refs, true)
}

// LinkAll replaces key words found in the root and it's
// children with links defined in the map. The map should be TEXT -> HREF
func linkAll(root *Element, refs map[string]string, caseSensitive bool) {
	WalkElements(root, func(e *Element) {
		for i, c := range e.Children {
			switch c := c.(type) {
			case string:
				for txt, href := range refs {
					txt := strings.ReplaceAll(txt, " ", `[\n\t\r\s]*`)
					exp := `\b(` + txt + `)\b`
					if !caseSensitive {
						exp = "(?i)" + exp
					}
					re := regexp.MustCompile(exp)
					c = re.ReplaceAllString(c, `<a href="`+href+`">$1</a>`)
				}
				e.Children[i] = c
			}
		}
	})
}

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
