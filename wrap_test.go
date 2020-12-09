package web

import (
	"bytes"
	"strings"
	"testing"
)

func TestWrap(t *testing.T) {
	var buf bytes.Buffer
	x := Wrap(Div("a"), Span(Class("c"), "s"))
	x.WriteTo(&buf)
	got := buf.String()
	if strings.Contains(got, "wrapper") {
		t.Error(got)
	}
}

func TestWrap_markdown(t *testing.T) {
	var buf bytes.Buffer
	NewFile("x.md", Section(H2("heading"), P("hepp"))).WriteTo(&buf)
	a := buf.String()
	buf.Reset()

	// with wrapped children
	NewFile("x.md", Section(Wrap(H2("heading"), P("hepp")))).WriteTo(&buf)
	b := buf.String()

	if a != b {
		t.Error(a, "\n----\n", b)
	}
}
