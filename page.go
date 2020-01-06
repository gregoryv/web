package web

import (
	"io"
	"os"
	"path"
)

func NewPage(filename string, el *Element) *Page {
	return &Page{
		Filename: filename,
		Element:  el,
	}
}

type Page struct {
	Filename string
	*Element
}

func (p *Page) SaveTo(dir string) error {
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
