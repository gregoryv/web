package site_test

import (
	"github.com/gregoryv/web/site"
	"testing"
)

func TestCheckAll(t *testing.T) {
	done := make(chan bool)
	broken := make(chan site.BrokenLink)
	var count int
	go func() {
		for _ = range broken {
			count++
		}
		done <- true
	}()
	site.CheckLinks("example/", broken)
	<-done
	expBroken := 1
	if expBroken != count  {
		t.Errorf("Expected %v broken links got %v", expBroken, count)
	}
}
