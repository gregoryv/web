package web

import (
	"io"

	"github.com/gregoryv/nexus"
)

func Html(children ...interface{}) *HtmlTag {
	return &HtmlTag{
		NewTag("html", children...),
	}
}

type HtmlTag struct {
	*Tag
}

func (h *HtmlTag) WriteTo(w io.Writer) (int64, error) {
	p, err := nexus.NewPrinter(w)
	p.Print("<!DOCTYPE html>\n\n")
	h.Tag.WriteTo(p)
	return p.Written, *err
}
