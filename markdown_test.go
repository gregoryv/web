package web

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestMarkdownEncoder_Encode(t *testing.T) {
	ok := func(el *Element, exp ...string) {
		t.Helper()
		var buf bytes.Buffer
		hw := NewMarkdownEncoder(&buf)
		hw.Encode(el)
		got := buf.String()
		assert := asserter.New(t)
		for _, exp := range exp {
			assert().Contains(got, exp)
		}
	}
	ok(Html(), ``)
	ok(Html(Body()), ``)
	ok(H1("x"), "# x")
	ok(Img(Src("x")), "![](x)")
	ok(Img(Alt("a"), Src("x")), "![a](x)")
	ok(Pre(`a
b`), `    a
    b`)
	ok(Pre("x"), "  x")
	ok(A("txt", Href("h")), "[txt](h)")

	// linked image
	el := A(
		Href("http://example.com"),
		Img(
			Src("http://mysite.com/image.png"),
			Alt("Me"),
		),
	)
	exp := "[![Me](http://mysite.com/image.png)](http://example.com)"
	ok(el, exp)

	// multiple linked image
	div := Div(el, el)
	exp = `[![Me](http://mysite.com/image.png)](http://example.com)
[![Me](http://mysite.com/image.png)](http://example.com)`
	ok(div, exp)

	ok(A(), "")

	ok(P(`hello
      friend`), `hello
friend`)

	ok(Ul(
		Li("a"),
		Li("b"),
	), `- a
- b`)
}
