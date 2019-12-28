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

func Head(c ...interface{}) *Tag { return NewTag("head", c...) }
func Body(c ...interface{}) *Tag { return NewTag("body", c...) }

func Meta(c ...interface{}) *Tag { return NewSimpleTag("meta", c...) }
func Img(c ...interface{}) *Tag  { return NewSimpleTag("img", c...) }

func Src(v string) *Attr     { return &Attr{name: "src", val: v} }
func Lang(v string) *Attr    { return &Attr{name: "lang", val: v} }
func Charset(v string) *Attr { return &Attr{name: "charset", val: v} }
func Name(v string) *Attr    { return &Attr{name: "name", val: v} }
func Content(v string) *Attr { return &Attr{name: "content", val: v} }
