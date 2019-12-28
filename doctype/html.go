package doctype

import (
	"fmt"
	"io"
)

func Html(children ...interface{}) *HtmlTag {
	return &HtmlTag{
		Tag{
			children: children,
			name:     "html",
		},
	}
}

type HtmlTag struct {
	Tag
}

func (h *HtmlTag) WriteTo(w io.Writer) {
	fmt.Fprint(w, "<!DOCTYPE html>\n\n")
	h.Tag.WriteTo(w)
}

func Body(children ...interface{}) *Tag {
	return &Tag{
		children: children,
		name:     "body",
	}
}
