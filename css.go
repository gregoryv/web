package web

import (
	"io"

	"github.com/gregoryv/nexus"
)

func NewCSS() *CSS {
	return &CSS{
		rules:   make([]*rule, 0),
		imports: make([]string, 0),
	}
}

type CSS struct {
	Media   string
	rules   []*rule
	imports []string
}

// Import
func (me *CSS) Import(url string) {
	me.imports = append(me.imports, url)
}

func (r *CSS) WriteTo(w io.Writer) (int64, error) {
	p, err := nexus.NewPrinter(w)
	if r.Media != "" {
		p.Print(r.Media)
		p.Println("{")
	}
	for _, imp := range r.imports {
		p.Printf("@import url('%s');\n", imp)
	}
	for _, rule := range r.rules {
		rule.WriteTo(p)
	}
	if r.Media != "" {
		p.Println("}")
	}
	return p.Written, *err
}

func (c *CSS) Style(selector string, propvals ...string) {
	c.rules = append(c.rules, &rule{selector, propvals})
}

type rule struct {
	selector string
	propvals []string
}

func (r *rule) WriteTo(w io.Writer) (int64, error) {
	p, err := nexus.NewPrinter(w)
	p.Print(r.selector)
	p.Println(" {")
	for _, pv := range r.propvals {
		p.Println(pv + ";")
	}
	p.Println("}")
	return p.Written, *err
}
