package toc

import (
	"regexp"
	"strings"

	"github.com/gregoryv/web"
)

// MakeTOC generates ids for all named elements and generates a TOC
// into the destination element.
func MakeTOC(dest, root *web.Element, names ...string) {
	GenerateIDs(root, names...)
	dest.With(ParseTOC(root, names...))
}

// ParseTOC returns ul > li of all named elements.
func ParseTOC(root *web.Element, names ...string) *web.Element {
	ul := web.Ul()
	web.WalkElements(root, func(e *web.Element) {
		for _, name := range names {
			if e.Name == name {
				a := web.A(web.Href("#"+idOf(e)), e.Text())
				ul.With(web.Li(web.Class(name), a))
				continue
			}
		}
	})
	return ul
}

func GenerateIDs(root *web.Element, names ...string) {
	web.WalkElements(root, func(e *web.Element) {
		if hasId(e) {
			return
		}
		for _, name := range names {
			if e.Name == name {
				e.With(web.Id(idOf(e)))
			}
		}
	})
}

var idChars = regexp.MustCompile(`\W`)

func idOf(e *web.Element) string {
	for _, attr := range e.Attributes {
		if attr.Name == "id" {
			return attr.Val
		}
	}
	txt := idChars.ReplaceAllString(e.Text(), "")
	return strings.ToLower(txt)
}

func hasId(e *web.Element) bool {
	for _, attr := range e.Attributes {
		if attr.Name == "id" {
			return true
		}
	}
	return false
}
