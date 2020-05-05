package web

import (
	"bytes"
	"fmt"
	"io"
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
}

func (t *Element) String() string {
	var buf bytes.Buffer
	t.WriteTo(&buf)
	return buf.String()
}

func (t *Element) WriteTo(w io.Writer) (int64, error) {
	return NewHtmlWriter(w).WriteHtml(t)
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

func (t *Element) hasAttr(name string) bool {
	for _, a := range t.Attributes {
		if a.Name == name {
			return true
		}
	}
	return false
}

func Attr(name string, val interface{}) *Attribute {
	return &Attribute{Name: name, Val: fmt.Sprintf("%v", val)}
}

type Attribute struct {
	Name string
	Val  string
}

func (a *Attribute) String() string {
	return fmt.Sprintf("%s=%q", a.Name, a.Val)
}
