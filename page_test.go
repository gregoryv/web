package web

import (
	"testing"

	"github.com/gregoryv/workdir"
)

func TestPage(t *testing.T) {
	wd, err := workdir.TempDir()
	if err != nil {
		t.Fatal(err)
	}
	defer wd.RemoveAll()
	page := NewPage("x.html", Html(Body()))
	page.SaveTo(string(wd))

	wd.Chmod("x.html", 0000)
	err = page.SaveTo(string(wd))
	if err == nil {
		t.Error("should fail")
	}
}
