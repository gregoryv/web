package web

import (
	"os"
	"testing"
)

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
