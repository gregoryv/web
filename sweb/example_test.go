package sweb_test

import (
	"os"

	. "github.com/gregoryv/web/sweb"
)

func Example() {
	Html()
	Body()
	Meta(Charset("utf-8"))

	Article()
	H1_("hello", Class("main-title"))
	Ul()
	Li("one")
	Li("two")
	EndArticle()

	Article(Class("final"))
	H1("world")

	html := End()
	html.WriteTo(os.Stdout)

	//output:
	//<!DOCTYPE html>
	//
	// <html>
	// <body>
	// <meta charset="utf-8"/>
	// <article>
	// <h1 class="main-title">hello</h1>
	// <ul>
	// <li>one<li>two</li>
	// </li>
	// </ul>
	// </article>
	// <article class="final">
	// <h1>world</h1>
	// </article>
	// </body>
	// </html>
}
