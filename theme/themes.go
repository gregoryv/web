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
		"line-height: 1.3em",
	)
	css.Style("a:link",
		"color: #0000EE",
	)
	css.Style("a:visited",
		"color: #551A8B",
	)
	return css
}

func GoishColors() *web.CSS {
	css := web.NewCSS()
	css.Style("a:link, a:visited",
		"color: #007d9c",
		"text-decoration: none",
	)
	css.Style("a:hover",
		"text-decoration: underline",
	)
	return css
}
