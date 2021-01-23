package web

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/workdir"
)

func Test_serving_page_over_http(t *testing.T) {
	var (
		p      = NewPage(Html(Body(H1("the secret lies within you"))))
		assert = asserter.New(t)
		exp    = assert().ResponseFrom(p)
	)
	exp.Contains("lies within you", "GET", "/")
	exp.Contains("<html>", "GET", "/")
}

func Test_save_page_without_filename(t *testing.T) {
	err := NewPage(Html()).SaveTo(".")
	if err == nil {
		t.Error("should fail")
	}
}

func TestPage(t *testing.T) {
	wd, err := workdir.TempDir()
	if err != nil {
		t.Fatal(err)
	}
	defer wd.RemoveAll()
	page := NewFile("x.md", Html(Body()))
	err = page.SaveTo(string(wd))
	if err != nil {
		t.Fatal(err)
	}

	wd.Chmod("x.html", 0000)
	err = page.SaveAs(wd.Join("/x.html"))
	if err == nil {
		t.Error("should fail")
	}
}

func TestPage_markdown(t *testing.T) {
	page := NewFile("x.md", Html(Body()))
	var md bytes.Buffer
	page.WriteTo(&md)

	page.Filename = "x.html"
	var html bytes.Buffer
	page.WriteTo(&html)

	if md.String() == html.String() {
		t.Error("expected markdown\n", md.String())
	}
}
