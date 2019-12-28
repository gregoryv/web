package doctype

import (
	"bytes"
	"testing"

	"github.com/gregoryv/golden"
)

func TestHtml(t *testing.T) {
	w := bytes.NewBufferString("")
	Html(
		Lang("en-US"),
		Head(
			Meta(Charset("utf-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
		),
		Body(
			"hello",
			Img(
				Src("img/example.png"),
			),
		),
	).WriteTo(w)
	golden.Assert(t, w.String())
}
