package web

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gregoryv/nexus"
)

func NewCSS() *CSS {
	return &CSS{
		rules:   make([]*rule, 0),
		imports: make([]string, 0),
	}
}

type CSS struct {
	media   string
	rules   []*rule
	imports []string

	medias []*CSS
}

// SetMedia sets the media and returns the same css
func (me *CSS) SetMedia(v string) {
	me.media = fmt.Sprintf("@media %s", v)
}

// Media adds a media section returning a new css for styling
func (me *CSS) Media(v string) *CSS {
	css := NewCSS()
	css.SetMedia(v)
	me.medias = append(me.medias, css)
	return css
}

// Import
func (me *CSS) Import(url string) {
	me.imports = append(me.imports, url)
}

func (me *CSS) WriteTo(w io.Writer) (int64, error) {
	p, err := nexus.NewPrinter(w)
	if me.media != "" {
		p.Print(me.media)
		p.Println("{")
	}
	for _, imp := range me.imports {
		p.Printf("@import url('%s');\n", imp)
	}
	for _, rule := range me.rules {
		rule.WriteTo(p)
	}
	if me.media != "" {
		p.Println("}")
	}
	for _, media := range me.medias {
		p.Println()
		media.WriteTo(w)
	}
	return p.Written, *err
}

func (me *CSS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/css")
	me.WriteTo(w)
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
