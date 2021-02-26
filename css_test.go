package web

import (
	"os"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestCSS_ServeHTTP(t *testing.T) {
	c := NewCSS()
	c.Style("#x", "margin: 0 0")

	assert := asserter.New(t)
	exp := assert().ResponseFrom(c)

	exp.StatusCode(200, "GET", "/")
	exp.Header("content-type", "text/css", "GET", "/")
}

func Test(t *testing.T) {
	c := NewCSS()
	c.Style("#x", "margin: 0 0")
}

func Example() {
	c := NewCSS()
	c.Import("https://fonts.googleapis.com/css?family=Open+Sans")
	c.Style("#x", "margin: 0 0")

	p := c.Media("print")
	p.Style("footer", "display: none")

	c.WriteTo(os.Stdout)
	// output:
	// @import url('https://fonts.googleapis.com/css?family=Open+Sans');
	// #x {
	// margin: 0 0;
	// }
	//
	// @media print{
	// footer {
	// display: none;
	// }
	// }
}
