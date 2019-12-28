package doctype_test

import (
	"bytes"
	"testing"

	"github.com/gregoryv/golden"
	. "github.com/gregoryv/web/doctype"
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
			Article(
				Section(
					H1("hello"),
					"world",
					Img(Src("img/example.png")),
				),
			),
		),
	).WriteTo(w)
	golden.Assert(t, w.String())
}
