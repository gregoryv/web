package site

import (
	"fmt"
	"testing"

	"github.com/gregoryv/asserter"
)

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
	CheckLinks("example/", broken)
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
