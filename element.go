package web

import (
	"fmt"
	"io"
	"strings"

	"github.com/gregoryv/nexus"
)

func NewElement(Name string, childOrAttr ...interface{}) *Element {
	tag := &Element{Name: Name}
	tag.With(childOrAttr...)
	return tag
}

func NewSimpleElement(Name string, childOrAttr ...interface{}) *Element {
	tag := &Element{Name: Name}
	tag.With(childOrAttr...)
	tag.simple = true
	return tag
}

type Element struct {
	Children   []interface{}
	Name       string
	Attributes []*Attribute

	// No closing tag, e.g. <br />
	simple bool
	// Indentation
	level int
}

func (t *Element) WriteTo(w io.Writer) (int64, error) {
	p, err := nexus.NewPrinter(w)
	t.open(p)
	var afterString bool
	for _, child := range t.Children {
		switch child := child.(type) {
		case *Element:
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

func (t *Element) open(p *nexus.Printer) {
	indent(p, t.level)
	p.Print("<", t.Name)
	for _, Attributes := range t.Attributes {
		p.Print(" ", Attributes.String())
	}
	if !t.simple {
		p.Print(">\n")
	}
}

func (t *Element) close(p *nexus.Printer) {
	if t.simple {
		p.Print("/>\n")
		return
	}
	p.Println()
	indent(p, t.level)
	p.Print("</", t.Name, ">\n")
}

func (t *Element) With(childOrAttr ...interface{}) *Element {
	t.fill(childOrAttr...)
	return t
}

func (t *Element) fill(childOrAttr ...interface{}) {
	for _, ca := range childOrAttr {
		switch ca := ca.(type) {
		case *Attribute:
			t.Attributes = append(t.Attributes, ca)
		default:
			t.Children = append(t.Children, ca)
		}
	}
}

type Attribute struct {
	Name string
	Val  string
}

func (a *Attribute) String() string {
	return fmt.Sprintf("%s=%q", a.Name, a.Val)
}

func (a *Attribute) WriteTo(w io.Writer) (int64, error) {
	n, err := fmt.Fprint(w, a.String())
	return int64(n), err
}
