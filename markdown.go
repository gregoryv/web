package web

import (
	"io"
	"strings"

	"github.com/gregoryv/nexus"
)

func NewMarkdownWriter(w io.Writer) *MarkdownWriter {
	p, err := nexus.NewPrinter(w)
	return &MarkdownWriter{
		Printer: p,
		err:     err,
	}
}

type MarkdownWriter struct {
	*nexus.Printer
	err    *error
	indent string // ie. for pre tags
}

func (p *MarkdownWriter) WriteMarkdown(e *Element) {
	p.writeElement(e)
}

func (p *MarkdownWriter) writeElement(t interface{}) {
	switch t := t.(type) {
	case *Element:
		p.open(t)
		for _, a := range t.Attributes {
			p.writeAttr(a)
		}
		for _, child := range t.Children {
			p.writeElement(child)
		}
		p.close(t)
	case string:
		if p.indent == "" || strings.Index(t, "\n") == -1 {
			p.Print(t)
			return
		}
		lines := strings.Split(t, "\n")
		for _, line := range lines {
			p.Print(p.indent, line, "\n")
		}
	}
}

var markdown = map[string]string{
	"h1": "# ",
	"h2": "## ",
	"h3": "### ",
	"h4": "#### ",
	"h5": "##### ",
	"h6": "###### ",
	"ul": "",
	"p":  "",
	"li": "- ",
	"hr": "----",
	"br": "\n",
}

func (p *MarkdownWriter) writeAttr(a *Attribute) {
	if a.Name == "src" {
		p.Printf("(%s)", a.Val)
	}
	if a.Name == "alt" {
		p.Printf("[%s]", a.Val)
	}

}

func (p *MarkdownWriter) open(t *Element) {
	switch t.Name {
	case "img":
		p.Print("!")
		if !t.hasAttr("alt") {
			p.Print("[]")
		}
	case "pre":
		p.indent = "  "
	default:
		p.Print(markdown[t.Name])
	}
}

func (p *MarkdownWriter) close(t *Element) {
	p.Println()
	switch t.Name {
	case "li", "span":
	case "pre":
		p.indent = ""
	default:
		p.Println()
	}
}
