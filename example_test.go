package web

import "os"

func ExamplePage_WriteTo() {
	article := Article(
		Class("fancy"),
		H1("Title of my article"),
	)
	page := NewPage("", Html(Body(article)))
	page.WriteTo(os.Stdout)

	// output:
	// <!DOCTYPE html>
	//
	// <html>
	// <body>
	// <article class="fancy">
	// <h1>Title of my article</h1>
	//
	// </article>
	// </body>
	// </html>
}
