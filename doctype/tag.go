package doctype

import (
	"fmt"
	"io"
)

type Tag struct {
	children []interface{}
	name     string
}

func (t *Tag) WriteTo(w io.Writer) error {
	fmt.Fprintf(w, "<%s>", t.name)
	for _, child := range t.children {
		switch child := child.(type) {
		case *Tag:
			child.WriteTo(w)
		case string:
			fmt.Fprintf(w, child)
		}
	}
	fmt.Fprintf(w, "</%s>", t.name)
	return nil
}
