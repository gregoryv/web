package web

import (
	"bytes"
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
		if e.Name == "a" { // inside an link already
			return
		}
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

// CheckLinks scans all .html files in the root directory for broken
// links. Only local href and src references are checked.
func CheckLinks(root string) error {
	c := make(chan BrokenLink)
	broken := make([]BrokenLink, 0)
	done := make(chan bool)
	go func() {
		for v := range c {
			broken = append(broken, v)
		}
		done <- true
	}()

	htmlFiles, _ := find.By(find.NewShellPattern("*.html"), root)
	for e := htmlFiles.Front(); e != nil; e = e.Next() {
		// Could be parallellized if needed here
		file, _ := e.Value.(string)
		fh, _ := os.Open(file)
		defer fh.Close()
		doc, _ := html.Parse(fh)
		checkLink(file, path.Dir(file), doc, c)
	}
	close(c)
	<-done
	return combinedError(broken)
}

func combinedError(broken []BrokenLink) error {
	if len(broken) > 0 {
		var buf bytes.Buffer
		for _, v := range broken {
			buf.WriteString(v.String())
			buf.WriteString("\n")
		}
		return fmt.Errorf("broken links found:\n%s", buf.String())
	}
	return nil
}

func checkLink(file, rel string, n *html.Node, broken chan BrokenLink) {
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
		checkLink(file, rel, c, broken)
	}
}
