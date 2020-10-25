package web

import (
	"testing"

	"github.com/gregoryv/workdir"
)

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
	page := NewFile("x.html", Html(Body()))
	page.SaveTo(string(wd))

	wd.Chmod("x.html", 0000)
	err = page.SaveAs(wd.Join("/x.html"))
	if err == nil {
		t.Error("should fail")
	}
}
