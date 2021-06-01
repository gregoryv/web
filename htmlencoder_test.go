package web

import (
	"bytes"
	"strings"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_HtmlEncoder_produces_safe_output(t *testing.T) {
	var buf bytes.Buffer
	h := NewSafeHtmlEncoder(&buf)
	h.Encode("<script>")
	h.Encode(B(Attr("<n>", "<v>")))
	got := buf.String()
	exp := `&lt;script&gt;<b &lt;n&gt;="&lt;v&gt;"></b>`
	if got != exp {
		t.Error("got", got)
		t.Error("exp", exp)
	}
}

func TestHtmlEncoder_Encode(t *testing.T) {
	ok := func(el *Element, exp ...string) {
		t.Helper()
		var buf bytes.Buffer
		hw := NewHtmlEncoder(&buf)
		hw.Encode(el)
		got := buf.String()
		assert := asserter.New(t)
		for _, exp := range exp {
			assert().Contains(got, exp)
		}
	}
	ok(Html(), `<html>
</html>`)
	ok(Html(Body(
		Header(),
		Nav(),
	)), `<html>
<body>
<header></header>
<nav></nav>
</body>
</html>`)

	ok(
		Html(Body(H1("hello"), "text")),
		`<html>
<body>
<h1>hello</h1>
text</body>
</html>`,
	)

	ok(Span(strings.NewReader("hello")),
		"<span>hello</span>",
	)
	ok(Span(1), "<span>1</span>")

	ok(
		A(Href("http://x.com"), "label"),
		`<a href="http://x.com">label</a>`,
	)

	ok(
		Article(
			Div(),
			H2(),
			P(),
		),
		`<article>
<div></div>
<h2></h2>
<p></p>
</article>`,
	)
}
