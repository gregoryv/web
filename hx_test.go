package web

import (
	"bytes"
	"strings"
	"testing"
)

func TestHn(t *testing.T) {
	ok := func(n *Hn, shouldContain ...string) {
		t.Helper()
		var buf bytes.Buffer
		myArticle(n).WriteTo(&buf)
		got := buf.String()
		for _, substr := range shouldContain {
			if !strings.Contains(got, substr) {
				t.Fatalf("using %#v\n\n%s\nIS MISSING %q", n, got, substr)
			}
		}
	}
	ok(NewHn(-2), "<h1>main", "<h6>last")
	ok(&Hn{}, "<h1>main", "<h6>last")
	ok(NewHn(2), "<h2>main", "<h7>last")
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
