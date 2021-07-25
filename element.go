package web

import (
	"bytes"
	"fmt"
	"html"
	"io"
	"strings"
)

func NewElement(name string, childOrAttr ...interface{}) *Element {
	tag := &Element{
		Name:     name,
		Children: make([]interface{}, 0, 5),
	}
	tag.With(childOrAttr...)
	return tag
}

func NewSimpleElement(name string, childOrAttr ...interface{}) *Element {
	tag := &Element{Name: name}
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
	enc := NewHtmlEncoder(w)
	enc.Encode(t)
	return enc.Written, *enc.err
}

func (t *Element) With(childOrAttr ...interface{}) *Element {
	t.fill(childOrAttr...)
	return t
}

func (t *Element) Text() string {
	return text(t)
}

func text(t interface{}) string {
	switch t := t.(type) {
	case *Element:
		parts := make([]string, 0)
		for _, c := range t.Children {
			parts = append(parts, text(c))
		}
		return strings.Join(parts, " ")
	default:
		return fmt.Sprintf("%v", t)
	}
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

// HasAttr returns true if the named attribute is found on the
// element.
func (t *Element) HasAttr(name string) bool {
	for _, a := range t.Attributes {
		if a.Name == name {
			return true
		}
	}
	return false
}

func (t *Element) hasChild(name string) bool {
	for _, c := range t.Children {
		if c, ok := c.(*Element); ok {
			if c.Name == name {
				return true
			}
		}
	}
	return false
}

// Attr returns the named attribute or nil.
func (t *Element) Attr(name string) *Attribute {
	for _, a := range t.Attributes {
		if a.Name == name {
			return a
		}
	}
	return nil
}

// AttrVal returns attribute value if it exists, empty string otherwise.
func (t *Element) AttrVal(name string) string {
	a := t.Attr(name)
	if a == nil {
		return ""
	}
	return a.Val
}

// Attr creates a new named attribute with the given value.
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

func (a *Attribute) SafeString() string {
	return fmt.Sprintf("%s=%q", html.EscapeString(a.Name), html.EscapeString(a.Val))
}

func WalkElements(root *Element, fn func(e *Element)) {
	fn(root)
	for _, c := range root.Children {
		if c, ok := c.(*Element); ok {
			WalkElements(c, fn)
		}
	}
}

// ----------------------------------------

func Comment(childOrAttr ...interface{}) *Element {
	return NewElement("!--", childOrAttr...)
}
