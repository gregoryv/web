package web

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// NewSafePage returns a page same as NewPage, only the output if
// written as html is escaped. See NewSafeHtmlEncoder constructor.
func NewSafePage(el *Element) *Page {
	return &Page{
		Element: el,
		safe:    true,
	}
}

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

	safe bool // use safe html encoder
}

// SaveAs sets filename and then save to the current directory.
func (me *Page) SaveAs(filename string) error {
	dir, file := filepath.Split(filename)
	if dir == "" {
		me.Filename = filename
		return me.SaveTo(".")
	}
	me.Filename = file
	return me.SaveTo(dir)
}

// SaveTo saves the page to the given directory. Fails if
// page.Filename is empty.
func (p *Page) SaveTo(dir string) error {
	if p.Filename == "" {
		return fmt.Errorf("SaveTo: missing filename")
	}
	os.MkdirAll(dir, 0722)
	w, err := os.Create(path.Join(dir, p.Filename))
	if err != nil {
		return err
	}
	p.WriteTo(w)
	w.Close()
	return nil
}

// Size returns the rendered size of the page in bytes. Note! the page
// is rendered once to count the bytes.
func (me *Page) Size() int {
	n, _ := me.WriteTo(ioutil.Discard)
	return int(n)
}

// WriteTo writes the page using the given writer. Page.Filename
// extension decides format.  .md for markdown, otherwise HTML.
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
		enc.safe = p.safe
		if p.Element != nil {
			enc.Encode(p.Element)
		}
		return enc.Written, *enc.err
	}
}

func (p *Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.WriteTo(w)
}

type encoder interface {
	Encode(t interface{}) error
}
