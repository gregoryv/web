package web

import (
	"fmt"
	"io"
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

func (p *Page) WriteTo(w io.Writer) (int64, error) {
	hw := NewHtmlWriter(w)
	hw.Print("<!DOCTYPE html>\n\n")
	if p.Element != nil {
		hw.WriteHtml(p.Element)
	}
	return hw.Written, *hw.err
}
