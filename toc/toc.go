package toc

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gregoryv/web"
)

// MakeTOC generates ids for all named elements and generates a TOC
// into the destination element returning the created UL element.
func MakeTOC(dest, root *web.Element, names ...string) *web.Element {
	GenerateIDs(root, names...)
	GenerateAnchors(root, names...)
	ul := ParseTOC(root, names...)
	dest.With(ul)
	return ul
}

// ParseTOC returns ul > li of all named elements.
func ParseTOC(root *web.Element, names ...string) *web.Element {
	ul := web.Ul()
	ids := make(cache)
	web.WalkElements(root, func(e *web.Element) {
		for _, name := range names {
			if e.Name == name {
				a := web.A(web.Href("#"+ids.idOf(e)), e.Text())
				ul.With(web.Li(web.Class(name), a))
				continue
			}
		}
	})
	return ul
}

func GenerateIDs(root *web.Element, names ...string) {
	ids := make(cache)
	web.WalkElements(root, func(e *web.Element) {
		if hasId(e) {
			return
		}
		for _, name := range names {
			if e.Name == name {
				newid := ids.idOf(e)
				e.With(web.Id(newid))
			}
		}
	})
}

func GenerateAnchors(root *web.Element, names ...string) {
	web.WalkElements(root, func(e *web.Element) {
		if !hasId(e) {
			return
		}
		for _, name := range names {
			if e.Name == name {
				a := web.A(web.Href("#" + e.AttrVal("id"))).With(e.Children...)
				e.Children = []interface{}{a}
			}
		}
	})
}

type cache map[string]int

func (me cache) idOf(e *web.Element) string {
	newid := idOf(e)
	n, found := me[newid]
	if found {
		newid = fmt.Sprintf("%s%d", newid, n)
	}
	me[newid]++
	return newid
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
