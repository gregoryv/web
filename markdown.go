package web

import (
	"io"
	"strings"

	"github.com/gregoryv/web/internal/nexus"
)

func NewMarkdownEncoder(w io.Writer) *MarkdownEncoder {
	p, err := nexus.NewPrinter(w)
	return &MarkdownEncoder{
		Printer: p,
		err:     err,
		trim:    true,
	}
}

type MarkdownEncoder struct {
	*nexus.Printer
	err    *error
	indent string // ie. for pre tags
	trim   bool   // pre tags should keep original whitespaces

	oneliner bool
}

func (p *MarkdownEncoder) Encode(t interface{}) error {
	p.writeElement(t)
	return *p.err
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
			if t.Name != "wrapper" {
				p.open(t)
			}
			for _, child := range t.Children {
				p.writeElement(child)
			}
			if t.Name != "wrapper" {
				p.close(t)
			}
		}
	case ElementBuilder:
		p.writeElement(t.BuildElement())
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
			if p.trim {
				line = strings.TrimSpace(line)
			}
			p.Print(p.indent, line, "\n")
		}
	}
}

func (p *MarkdownEncoder) open(t *Element) {
	switch t.Name {
	case "pre":
		p.indent = "    "
		p.trim = false
	default:
		// todo: default to writing html
		p.Print(markdown[t.Name])
	}

}

func (p *MarkdownEncoder) close(t *Element) {
	switch t.Name {
	case "span", "body", "html":
	case "pre":
		p.Println()
		p.indent = ""
		p.trim = true
	case "h1", "h2", "h3", "h4", "h5", "h6":
		p.Println()
		p.Println()
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
