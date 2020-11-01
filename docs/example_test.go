package docs

import . "github.com/gregoryv/web"

func NewHtmlPage() *Page {
	page := NewPage(
		Html(
			Head(
				Title("my home page"),
			),
			Body(
				H1("My home"),
				"some text here",
			),
		),
	)
	return page
}
