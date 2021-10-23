package web

import (
	"os"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestCSS_SaveAs(t *testing.T) {
	c := NewCSS()
	c.Style("body", "padding: 0 0")

	out := "out.css"
	defer os.RemoveAll(out)

	err := c.SaveAs(out)
	if err != nil {
		t.Fatal(err)
	}

	os.Chmod(out, 0000)
	err = c.SaveAs(out)
	if err == nil {
		t.Error("should fail")
	}
}

func TestCSS_SaveTo_fails(t *testing.T) {
	c := NewCSS()
	if err := c.SaveTo("."); err == nil {
		t.Error("Filename not set, should fail")
	}
}

func TestCSS_ServeHTTP(t *testing.T) {
	c := NewCSS()
	c.Style("#x", "margin: 0 0")

	assert := asserter.New(t)
	exp := assert().ResponseFrom(c)

	exp.StatusCode(200, "GET", "/")
	exp.Header("content-type", "text/css", "GET", "/")
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

func ExampleCSS_With() {
	a := NewCSS()
	a.Style("body", "margin: 0 0")

	b := NewCSS()
	b.Style("p", "color:red")

	a.With(b)

	a.WriteTo(os.Stdout)
	// output:
	// body {
	// margin: 0 0;
	// }
	// p {
	// color:red;
	// }
}
