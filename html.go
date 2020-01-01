package web

import (
	"io"

	"github.com/gregoryv/nexus"
)

func Html(Children ...interface{}) *HtmlTag {
	return &HtmlTag{
		NewElement("html", Children...),
	}
}

type HtmlTag struct {
	*Element
}

func (h *HtmlTag) WriteTo(w io.Writer) (int64, error) {
	p, err := nexus.NewPrinter(w)
	p.Print("<!DOCTYPE html>\n\n")
	h.Element.WriteTo(p)
	return p.Written, *err
}
