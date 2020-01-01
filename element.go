package web

import (
	"fmt"
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
