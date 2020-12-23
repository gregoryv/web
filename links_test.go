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
		P(`Hello world at example.com`),
	)
	refs := map[string]string{
		"hello world": "http://example.com",
	}
	LinkAll(root, refs)
	var buf bytes.Buffer
	root.WriteTo(&buf)

	got := buf.String()
	if !strings.Contains(got, "http://example.com") {
		t.Error(got)
	}
	if !strings.Contains(got, ">Hello world<") {
		t.Error(got)
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
