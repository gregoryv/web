package doctype

import (
	"fmt"
	"io"

	"github.com/gregoryv/nexus"
)

type Tag struct {
	children []interface{}
	name     string
	attr     []interface{}
}

func (t *Tag) WriteTo(w io.Writer) (int, error) {
	p, err := nexus.NewPrinter(w)
	t.OpenTo(p)
	for _, child := range t.children {
		switch child := child.(type) {
		case *Tag:
			child.WriteTo(p)
		case string:
			p.Print(child)
		}
	}
	p.Printf("</%s>", t.name)
	return p.Written, *err
}

func (t *Tag) OpenTo(w io.Writer) (int, error) {
	return fmt.Fprint(w, "<", t.name, ">")
}
