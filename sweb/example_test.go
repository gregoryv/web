package sweb_test

import (
	"os"

	. "github.com/gregoryv/web"
	. "github.com/gregoryv/web/sweb"
)

func Example() {
	Html_(
		Head(
			Meta(Charset("utf-8")),
		),
	)
	Body_()

	Article_(H1("hello", Class("main-title")))

	Ul_(
		Li("one"),
		Li("two"),
	)
	EndUl()

	Article_(Class("final"))
	H1_("world")

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
