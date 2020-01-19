package web

import (
	"io"
	"strings"

	"github.com/gregoryv/nexus"
)

func NewHtmlWriter(w io.Writer) *HtmlWriter {
	p, err := nexus.NewPrinter(w)
	return &HtmlWriter{
		Printer: p,
		err:     err,
	}
}

type HtmlWriter struct {
	*nexus.Printer
	err *error
}

func (p *HtmlWriter) WriteHtml(t interface{}) (int64, error) {
	p.writeElement(t)
	return p.Written, *p.err
}

func (p *HtmlWriter) writeElement(t interface{}) {
	switch t := t.(type) {
	case *Element:
		p.open(t)
		for _, child := range t.Children {
			p.writeElement(child)
		}
		p.close(t)
	case io.Reader:
		io.Copy(p, t)
	default:
		p.Printf("%v", t)
	}
}

func (p *HtmlWriter) open(t *Element) {
	p.Print("<", t.Name)
	for _, Attributes := range t.Attributes {
		p.Print(" ", Attributes.String())
	}
	if !t.simple {
		p.Print(">")
	}
	if strings.Index("html body head article section ol ul dl", t.Name) > -1 {
		p.Println()
	}
}

func (p *HtmlWriter) close(t *Element) {
	if t.simple {
		p.Print("/>")
		if strings.Index("meta link", t.Name) > -1 {
			p.Println()
		}
		return
	}
	if strings.Index("html body head article section style script", t.Name) > -1 {
		p.Println()
	}
	p.Print("</", t.Name, ">")
	if strings.Index("p h1 h2 h3 h4 h5 h6 li ul ol dt dd head title style script", t.Name) > -1 {
		p.Println()
	}
}
