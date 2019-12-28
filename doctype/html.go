package doctype

import (
	"fmt"
	"io"
)

func Html(children ...interface{}) *HtmlTag {
	return &HtmlTag{
		NewTag("html", children...),
	}
}

type HtmlTag struct {
	*Tag
}

func (h *HtmlTag) WriteTo(w io.Writer) {
	fmt.Fprint(w, "<!DOCTYPE html>\n\n")
	h.Tag.WriteTo(w)
}
