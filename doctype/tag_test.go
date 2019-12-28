package doctype

import (
	"bytes"
	"testing"

	"github.com/gregoryv/golden"
)

func TestHtml(t *testing.T) {
	html := Html(
		Body(
			"hello",
		),
	)
	w := bytes.NewBufferString("")
	html.WriteTo(w)
	golden.Assert(t, w.String())
}
