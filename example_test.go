package web

import "os"

func ExamplePage_WriteTo() {
	article := Article(
		Class("fancy"),
		H1("Title of my article"),
		P(
			"before ", A(Href("http://example.com"), "example"),
			" after",
		),
	)
	page := NewPage(Html(Body(article)))
	page.WriteTo(os.Stdout)

	// output:
	// <!DOCTYPE html>
	//
	// <html>
	// <body>
	// <article class="fancy">
	// <h1>Title of my article</h1>
	// <p>before <a href="http://example.com">example</a> after</p>
	// </article>
	// </body>
	// </html>
}
