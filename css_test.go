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
	c.Media = "@media screen"
	c.Style("#x", "margin: 0 0")
	c.WriteTo(os.Stdout)
	// output:
	// @media screen{
	// #x {
	// margin: 0 0;
	// }
	// }
}
