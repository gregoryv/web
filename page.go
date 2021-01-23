package web

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// NewPage returns a page ready to be rendered. Filename is empty and
// must be set before saving.
func NewPage(el *Element) *Page {
	return &Page{
		Element: el,
	}
}

// NewFile returns a page with filename set, ready to be saved.
func NewFile(filename string, el *Element) *Page {
	return &Page{
		Filename: filename,
		Element:  el,
	}
}

type Page struct {
	Filename string
	*Element
}

// SaveAs sets filename and then save to the current directory.
func (me *Page) SaveAs(filename string) error {
	me.Filename = filename
	return me.SaveTo(".")
}

// SaveTo saves the page to the given directory. Fails if
// page.Filename is empty.
func (p *Page) SaveTo(dir string) error {
	if p.Filename == "" {
		return fmt.Errorf("page SaveTo: missing filename")
	}
	w, err := os.Create(path.Join(dir, p.Filename))
	if err != nil {
		return err
	}
	p.WriteTo(w)
	w.Close()
	return nil
}

// WriteTo writes the page using the given writer. Page.Filename
// extension decides format.  .md for markdown, otherwise HTML.
// markdown, html otherwise.
func (p *Page) WriteTo(w io.Writer) (int64, error) {
	switch path.Ext(p.Filename) {
	case ".md":
		enc := NewMarkdownEncoder(w)
		if p.Element != nil {
			enc.Encode(p.Element)
		}
		return enc.Written, *enc.err
	default:
		enc := NewHtmlEncoder(w)
		if p.Element != nil {
			enc.Encode(p.Element)
		}
		return enc.Written, *enc.err
	}
}

func (p *Page) ServeAs(filename string) http.HandlerFunc {
	p.Filename = filename
	return func(w http.ResponseWriter, _ *http.Request) {
		p.WriteTo(w)
	}
}

type encoder interface {
	Encode(t interface{}) error
}
