package web

import (
	"io"

	"github.com/gregoryv/nexus"
)

func NewHtmlEncoder(w io.Writer) *HtmlEncoder {
	p, err := nexus.NewPrinter(w)
	return &HtmlEncoder{
		Printer: p,
		err:     err,
	}
}

type HtmlEncoder struct {
	*nexus.Printer
	err *error
}

func (p *HtmlEncoder) Encode(t interface{}) (int64, error) {
	if t, ok := t.(*Element); ok && t.Name == "html" {
		p.Print("<!DOCTYPE html>\n\n")
	}
	p.writeElement(t)
	return p.Written, *p.err
}

func (p *HtmlEncoder) writeElement(t interface{}) {
	switch t := t.(type) {
	case *Element:
		p.open(t)
		for _, child := range t.Children {
			p.writeElement(child)
		}
		p.close(t)
	case io.Reader:
		io.Copy(p, t)
	case io.WriterTo:
		t.WriteTo(p)
	default:
		p.Printf("%v", t)
	}
}

var newLineAfterOpen = map[string]bool{
	"html":     true,
	"body":     true,
	"head":     true,
	"table":    true,
	"article":  true,
	"section":  true,
	"ol":       true,
	"ul":       true,
	"dl":       true,
	"fieldset": true,
}

func (p *HtmlEncoder) open(t *Element) {
	p.Print("<", t.Name)
	for _, Attributes := range t.Attributes {
		p.Print(" ", Attributes.String())
	}
	if !t.simple {
		p.Print(">")
	}
	if newLineAfterOpen[t.Name] {
		p.Println()
	}
}

var newLineAfterClose = map[string]bool{
	"html":     true,
	"body":     true,
	"head":     true,
	"title":    true,
	"table":    true,
	"tr":       true,
	"article":  true,
	"header":   true,
	"nav":      true,
	"section":  true,
	"style":    true,
	"script":   true,
	"meta":     true,
	"link":     true,
	"p":        true,
	"h1":       true,
	"h2":       true,
	"h3":       true,
	"h4":       true,
	"h5":       true,
	"h6":       true,
	"li":       true,
	"ul":       true,
	"ol":       true,
	"dt":       true,
	"dd":       true,
	"div":      true,
	"fieldset": true,
}

func (p *HtmlEncoder) close(t *Element) {
	if t.simple {
		p.Print("/>")
		if newLineAfterClose[t.Name] {
			p.Println()
		}
		return
	}
	p.Print("</", t.Name, ">")
	if newLineAfterClose[t.Name] {
		p.Println()
	}
}
