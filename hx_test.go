package web

import (
	"bytes"
	"strings"
	"testing"
)

func Test_hx(t *testing.T) {
	var (
		page = NewPage(myArticle(&Hn{2}))
		buf  bytes.Buffer
	)
	page.WriteTo(&buf)
	got := buf.String()
	if strings.Contains(got, "<h1>") {
		t.Error(got)
	}
	if !strings.Contains(got, "<h2>main") {
		t.Error(got)
	}
	if !strings.Contains(got, "<h8>last") {
		t.Error(got)
	}

}

func myArticle(n *Hn) *Element {
	return Article(
		n.H1("main"),
		P(""),
		Section(
			n.H2("next"),
		),
		n.H3("sub"),
		n.H4("sub"),
		n.H5("sub"),
		n.H6("last"),
	)
}
