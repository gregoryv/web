package docs

import (
	"bytes"
	"html"

	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/files"
	"github.com/gregoryv/web/toc"
)

func NewIndexPage() *Page {
	page := NewPage(
		Html(
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				Title("gregoryv/web"),
				Style(Theme()),
			),
			Body(
				IndexArticle(),
			),
		),
	)
	return page
}

func IndexArticle() *Element {
	nav := Nav(P("Table of contents"))
	var example bytes.Buffer
	NewHtmlPage().WriteTo(&example)

	article := Article(
		H1("Web - Go module for HTML generation"),

		P(`This package provides a straight forward way of programming
		   HTML, leaving the separation of view and model as a design
		   choice for the developer. Also no values are automatically
		   escaped.`),

		nav,

		H2("Install"),
		Code(
			"go get ", A(
				Href("https://github.com/gregoryv/web"),
				"github.com/gregoryv/web/...",
			),
		),
		H2("API documentation"),
		Ul(
			Li(
				A(
					Href("https://godoc.org/github.com/gregoryv/web"),
					"web",
				), " - HTML generation",
			),
			Li(
				A(
					Href("https://godoc.org/github.com/gregoryv/web/files"),
					"web/files",
				), " - file loading tools",
			),
			Li(
				A(
					Href("https://godoc.org/github.com/gregoryv/web/apidoc"),
					"web/apidoc",
				), " - HTTP request/response documentation tools",
			),
			Li(
				A(
					Href("https://godoc.org/github.com/gregoryv/web/toc"),
					"web/toc",
				), " - table of content tools",
			),
		),
		H2("Example"),
		Pre(
			Code(Class("srcfile"),
				files.MustLoad("example_test.go"),
			),
		),
		P("Produces"),
		Pre(
			Code(Class("srcfile output"),
				html.EscapeString(example.String()),
			),
		),

		H2("About"),
		P("Written by ", gregory, Br(),
			"MIT License",
		),

		H2("License"),
		Pre(
			files.MustLoad("../LICENSE"),
		),
	)
	toc.MakeTOC(nav, article, "h2", "h3")
	return article
}

const gregory = "Gregory Vin&ccaron;i&cacute;"
