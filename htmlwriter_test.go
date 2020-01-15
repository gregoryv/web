package web

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestHtmlWriter_WriteHtml(t *testing.T) {
	ok := func(el *Element, exp ...string) {
		t.Helper()
		var buf bytes.Buffer
		hw := NewHtmlWriter(&buf)
		hw.WriteHtml(el)
		got := buf.String()
		assert := asserter.New(t)
		for _, exp := range exp {
			assert().Contains(got, exp)
		}
	}
	ok(Html(), `<html>

</html>`)
	ok(Html(Body()), `<html>
<body>

</body>
</html>`)

	ok(
		Html(Body(H1("hello"), "text")),
		`<html>
<body>
<h1>hello</h1>
text
</body>
</html>`,
	)

}