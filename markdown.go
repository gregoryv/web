package web

import (
	"io"
	"strings"

	"github.com/gregoryv/nexus"
)

func NewMarkdownEncoder(w io.Writer) *MarkdownEncoder {
	p, err := nexus.NewPrinter(w)
	return &MarkdownEncoder{
		Printer: p,
		err:     err,
	}
}

type MarkdownEncoder struct {
	*nexus.Printer
	err    *error
	indent string // ie. for pre tags

	oneliner bool
}

func (p *MarkdownEncoder) WriteMarkdown(e *Element) {
	p.writeElement(e)
}

func (p *MarkdownEncoder) writeElement(t interface{}) {
	switch t := t.(type) {
	case *Element:
		switch t.Name {
		case "img":
			p.Print("![", t.AttrVal("alt"), "](", t.AttrVal("src"), ")")
			if p.oneliner {
				return
			}
			p.Print("\n")
		case "a":
			p.Print("[")
			for _, child := range t.Children {
				p.oneliner = true
				p.writeElement(child)
				p.oneliner = false
			}
			p.Print("]")
			href := t.Attr("href")
			if href == nil {
				return
			}
			p.Print("(", href.Val, ")")
			if t.hasChild("img") { // special case to put linked images on separate lines
				p.Print("\n")
			}
		default:
			p.open(t)
			for _, child := range t.Children {
				p.writeElement(child)
			}
			for _, a := range t.Attributes {
				p.writeAttr(a)
			}
			p.close(t)
		}
	case string:
		if p.oneliner {
			p.Print(t)
			return
		}
		if strings.Index(t, "\n") == -1 {
			p.Print(p.indent, t)
			return
		}
		lines := strings.Split(t, "\n")
		for _, line := range lines {
			p.Print(p.indent, line, "\n")
		}
	}
}

// printAttr
func (p *MarkdownEncoder) printAttr(attr *Attribute) {
	if attr == nil {
		return
	}
	p.Print(attr.Val)
}

// printOneline
func (p *MarkdownEncoder) printOneline(t []interface{}) {
	for _, t := range t {
		switch t := t.(type) {
		case string:
			p.Print(t)
		}
	}
}

func (p *MarkdownEncoder) open(t *Element) {
	switch t.Name {
	case "pre":
		p.indent = "    "
	default:
		p.Print(markdown[t.Name])
	}
}

func (p *MarkdownEncoder) writeAttr(a *Attribute) {
	switch a.Name {
	case "src":
		p.Printf("(%s)", a.Val)
	case "alt":
		p.Printf("[%s]", a.Val)
	}
}

func (p *MarkdownEncoder) close(t *Element) {
	p.Println()
	switch t.Name {
	case "li", "span":
	case "pre":
		p.indent = ""
	default:
		p.Println()
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
