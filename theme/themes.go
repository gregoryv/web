// Package theme provides some basic web page themes
package theme

import "github.com/gregoryv/web"

func GoldenSpace() *web.CSS {
	css := web.NewCSS()
	css.Style("html, body",
		"margin: 0 0",
		"padding: 0 0",
	)
	css.Style("body",
		"padding: 1em 1.612em",
	)
	css.Style("p, li",
		"line-height: 1.4em",
	)
	return css
}

func GoishColors() *web.CSS {
	css := web.NewCSS()
	css.Style("a",
		"color: #007d9c",
		"text-decoration: none",
	)
	css.Style("a:hover",
		"text-decoration: underline",
	)
	return css
}
