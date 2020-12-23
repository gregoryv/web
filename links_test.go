package web

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestLinkAll(t *testing.T) {
	root := Article(
		Section(
			P(`Larnic gazed at the stars, surrounded by the black
               mountains.`),
		),
	)
	refs := map[string]string{
		"larnic":          "http://example.com",
		"black mountains": "http://black.com",
	}
	LinkAll(root, refs)

	var buf bytes.Buffer
	root.WriteTo(&buf)

	got := buf.String()
	ok := strings.Contains(got, "http://example.com") &&
		strings.Contains(got, ">Larnic") &&
		strings.Contains(got, "Larnic<")
	if !ok {
		t.Error(got)
	}

	ok = strings.Contains(got, ">black") &&
		strings.Contains(got, "mountains<")
	if !ok {
		t.Error("multiline link failed\n", got)
	}
}

func TestCheckLinks(t *testing.T) {
	done := make(chan bool)
	broken := make(chan BrokenLink)
	var count int
	go func() {
		for _ = range broken {
			count++
		}
		done <- true
	}()
	CheckLinks("./internal/example/", broken)
	<-done
	assert := asserter.New(t)
	assert().Equals(count, 1)
}

func TestBrokenLink(t *testing.T) {
	a := BrokenLink{"a", "a", fmt.Errorf("err")}
	b := BrokenLink{"a", "b", nil}

	assert := asserter.New(t)
	assert(a.String() != b.String()).Error("String() is same for a and b")
	assert().Contains(a.String(), "err")
}
