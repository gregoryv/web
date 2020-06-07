package toc

import (
	"regexp"
	"strings"

	"github.com/gregoryv/web"
)

func GenerateIds(e *web.Element, names ...string) {
	for _, child := range e.Children {
		switch child := child.(type) {
		case *web.Element:
			for _, name := range names {
				if child.Name == name {
					// fix id

					txt := contentText(child)
					id := strings.ToLower(txt)
					child.With(web.Id(id))
				}
			}
		}
	}
}

var idChars = regexp.MustCompile(`\W`)

func contentText(e *web.Element) string {
	return idChars.ReplaceAllString(e.Text(), "")
}
