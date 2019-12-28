package doctype

import (
	"fmt"
	"io"

	"github.com/gregoryv/nexus"
)

func NewTag(name string, childOrAttr ...interface{}) *Tag {
	tag := &Tag{name: name}
	tag.fill(childOrAttr...)
	return tag
}

type Tag struct {
	children []interface{}
	name     string
	attr     []*Attr

	// No closing tag, e.g. <br />
	simple bool
}

func (t *Tag) WriteTo(w io.Writer) (int, error) {
	p, err := nexus.NewPrinter(w)
	t.open(p)
	for _, child := range t.children {
		switch child := child.(type) {
		case *Tag:
			child.WriteTo(p)
		case string:
			p.Print(child)
		}
	}
	t.close(p)
	return p.Written, *err
}

func (t *Tag) open(p *nexus.Printer) {
	p.Print("<", t.name)
	for _, attr := range t.attr {
		p.Print(" ", attr.String())
	}
	if !t.simple {
		p.Print(">")
	}
}

func (t *Tag) close(p *nexus.Printer) {
	if t.simple {
		p.Print("/>")
		return
	}
	p.Print("</", t.name, ">")
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
