package web

import (
	"io"
	"os"
	"path"
)

func NewPage(filename string, content io.WriterTo) *Page {
	return &Page{
		filename: filename,
		WriterTo: content,
	}
}

type Page struct {
	filename string
	io.WriterTo
}

func (p *Page) SaveTo(dir string) error {
	fh, err := os.Create(path.Join(dir, p.filename))
	if err != nil {
		return err
	}
	p.WriteTo(fh)
	fh.Close()
	return nil
}
