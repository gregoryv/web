package doctype

import (
	"fmt"
	"io"
	"strings"

	"github.com/gregoryv/nexus"
)

func NewTag(name string, childOrAttr ...interface{}) *Tag {
	tag := &Tag{name: name}
	tag.With(childOrAttr...)
	return tag
}

func NewSimpleTag(name string, childOrAttr ...interface{}) *Tag {
	tag := &Tag{name: name}
	tag.With(childOrAttr...)
	tag.simple = true
	return tag
}

type Tag struct {
	children []interface{}
	name     string
	attr     []*Attr

	// No closing tag, e.g. <br />
	simple bool
	// Indentation
	level int
}

func (t *Tag) WriteTo(w io.Writer) (int, error) {
	p, err := nexus.NewPrinter(w)
	t.open(p)
	var afterString bool
	for _, child := range t.children {
		switch child := child.(type) {
		case *Tag:
			child.level = t.level + 1
			if afterString {
				child.level = 0
				p.Print(" ")
			}
			child.WriteTo(p)
			afterString = false
		case string:
			indent(p, t.level+1)
			p.Print(child)
			afterString = true
		}
	}
	t.close(p)
	return p.Written, *err
}

func indent(p *nexus.Printer, level int) {
	p.Print(strings.Repeat("  ", level))
}

func (t *Tag) open(p *nexus.Printer) {
	indent(p, t.level)
	p.Print("<", t.name)
	for _, attr := range t.attr {
		p.Print(" ", attr.String())
	}
	if !t.simple {
		p.Print(">\n")
	}
}

func (t *Tag) close(p *nexus.Printer) {
	if t.simple {
		p.Print("/>\n")
		return
	}
	p.Println()
	indent(p, t.level)
	p.Print("</", t.name, ">\n")
}

func (t *Tag) With(childOrAttr ...interface{}) {
	t.fill(childOrAttr...)
}

func (t *Tag) fill(childOrAttr ...interface{}) {
	for _, ca := range childOrAttr {
		switch ca := ca.(type) {
		case *Attr:
			t.attr = append(t.attr, ca)
		default:
			t.children = append(t.children, ca)
		}
	}
}

type Attr struct {
	name string
	val  string
}

func (a *Attr) String() string {
	return fmt.Sprintf("%s=%q", a.name, a.val)
}

func (a *Attr) WriteTo(w io.Writer) (int, error) {
	return fmt.Fprint(w, a.String())
}
