package web

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/gregoryv/asserter"
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

func TestCheckLinks_fails(t *testing.T) {
	err := CheckLinks("./testdata/")
	if err == nil {
		t.Error("didn't")
	}
}

func TestBrokenLink(t *testing.T) {
	a := BrokenLink{"a", "a", fmt.Errorf("err")}
	b := BrokenLink{"a", "b", nil}

	assert := asserter.New(t)
	assert(a.String() != b.String()).Error("String() is same for a and b")
	assert().Contains(a.String(), "err")
}

func Test_combinedError(t *testing.T) {
	err := combinedError(make([]BrokenLink, 0))
	if err != nil {
		t.Fatal(err)
	}
}
