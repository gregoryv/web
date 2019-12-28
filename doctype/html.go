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

func Body(children ...interface{}) *Tag {
	return NewTag("body", children...)
}

func Img(children ...interface{}) *Tag {
	tag := NewTag("img", children...)
	tag.simple = true
	return tag
}

func Src(val string) *Attr {
	return &Attr{
		name: "src",
		val:  val,
	}
}
