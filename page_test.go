package web

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_safe_page_over_http(t *testing.T) {
	var (
		p      = NewSafePage(Html(Body(H1("<script>"))))
		assert = asserter.New(t)
		exp    = assert().ResponseFrom(p)
	)
	exp.Contains("&lt;script&gt;", "GET", "/")
}

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

func TestPage_Size(t *testing.T) {
	page := newExamplePage()
	got := page.Size()
	if got != 1024 {
		t.Error("unexpected size", got)
	}
}

func TestPage(t *testing.T) {
	page := NewFile("x.md", Html(Body()))

	t.Run("SaveTo", func(t *testing.T) {
		if err := page.SaveTo(t.TempDir()); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("SaveAs with insufficient permissions", func(t *testing.T) {
		out := filepath.Join(t.TempDir(), "x.html")
		page.SaveAs(out)    // save the page first
		os.Chmod(out, 0000) // remove permissions to change it
		if err := page.SaveAs(out); err == nil {
			t.Error("should fail")
		}
	})

	t.Run("SaveAs current directory", func(t *testing.T) {
		os.Chdir(t.TempDir()) // switch the current directory
		if err := page.SaveAs("xa.html"); err != nil {
			t.Error(err)
		}
		// check the file exists with the correct name
		if _, err := os.Open("xa.html"); err != nil {
			t.Error(err)
		}
	})

	t.Run("SaveAs current directory", func(t *testing.T) {
		os.Chdir(t.TempDir()) // switch the current directory
		if err := page.SaveAs("a/b/c.html"); err != nil {
			t.Error(err)
		}
		// check the file exists with the correct name and path
		if _, err := os.Open("a/b/c.html"); err != nil {
			t.Error(err)
		}
	})

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
