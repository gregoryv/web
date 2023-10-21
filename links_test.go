package web

import (
	"bytes"
	"strings"
	"testing"
)

func TestILinkAll(t *testing.T) {
	root := Article(
		Section(
			P(`Larnic gazed at the stars, surrounded by the black
               mountains.`),
			A(Href("/"), "larnic"),
		),
	)
	refs := map[string]string{
		"larnic":          "http://example.com",
		"black mountains": "http://black.com",
	}
	ILinkAll(root, refs)

	var buf bytes.Buffer
	root.WriteTo(&buf)

	got := buf.String()
	ok := strings.Contains(got, "http://example.com") &&
		strings.Contains(got, ">Larnic") &&
		strings.Contains(got, "Larnic<") &&
		!strings.Contains(got, "</a></a>")
	if !ok {
		t.Error(got)
	}

	ok = strings.Contains(got, ">black") &&
		strings.Contains(got, "mountains<")
	if !ok {
		t.Error("multiline link failed\n", got)
	}
}

func TestLinkAll(t *testing.T) {
	root := Article(
		Section(
			P(`One car too many.`),
		),
	)
	refs := map[string]string{
		"ne": "#ne",
	}
	LinkAll(root, refs)

	var buf bytes.Buffer
	root.WriteTo(&buf)

	got := buf.String()
	if strings.Contains(got, "#ne") {
		t.Error(got)
	}
}
