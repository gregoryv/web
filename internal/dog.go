package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/gregoryv/nexus"
)

//go:generate go run .
func main() {
	generateWeb()
}

func generateWeb() {
	var buf bytes.Buffer
	p, _ := nexus.NewPrinter(&buf)
	p.Println("package web")
	p.Println("// Code generated by internal/dog.go DO NOT EDIT!")
	p.Println()
	// With children
	writeElements(p, complexElements...)
	p.Println()

	writeSimpleElements(p, simpleElements...)
	p.Println()

	writeAttributes(p, attributes...)

	// write result
	w, _ := os.Create("../generated.go")
	io.Copy(w, &buf)
	w.Close()

	// tidy output
	out, err := exec.Command("gofmt", "-w", w.Name()).CombinedOutput()
	if err != nil {
		fmt.Println(err, string(out))
		os.Exit(1)
	}
}

var complexElements = []string{
	"a",
	"abbr",
	"acronym",
	"address",
	"article",
	"aside",
	"b",
	"big",
	"blockquote",
	"body",
	"button",
	"cite",
	"code",
	"dd",
	"del",
	"details",
	"dfn",
	"div",
	"dl",
	"dt",
	"em",
	"footer",
	"form",
	"fieldset",
	"h1",
	"h2",
	"h3",
	"h4",
	"h5",
	"h6",
	"head",
	"header",
	"hgroup",
	"html",
	"i",
	"ins",
	"kbd",
	"label",
	"legend",
	"li",
	"mark",
	"menu",
	"meter",
	"nav",
	"noscript",
	"ol",
	"optgroup",
	"option",
	"output",
	"p",
	"pre",
	"quote",
	"script",
	"section",
	"select",
	"span",
	"style",
	"sub",
	"summary",
	"sup",
	"table",
	"tbody",
	"td",
	"textarea",
	"th",
	"thead",
	"title",
	"tr",
	"u",
	"ul",
	"var",
}

var simpleElements = []string{
	"base",
	"br",
	"hr",
	"img",
	"input",
	"keygen",
	"link",
	"meta",
}

var attributes = []string{
	"action",
	"alt",
	"autocomplete",
	"charset",
	"class",
	"content",
	"href",
	"id",
	"lang",
	"method",
	"name",
	"rel",
	"src",
	"tabindex",
	"type",
	"value",

	"onBlur",
	"onFocus",
	"onLoad",
	"onMouseDown",
	"onMouseOut",
	"onMouseOver",
	"onMouseUp",
	"onMouseWheel",
}

// ----------------------------------------

func writeElements(p *nexus.Printer, tags ...string) {
	for _, Name := range tags {
		funcName := capitalize(Name)
		p.Printf(
			`// %s returns an <%s> element with optional children or attributes
func %s(c ...interface{}) *Element {
    return NewElement("%s", c...)
}`,
			funcName, Name, funcName, Name,
		)
		p.Println()
	}
}

func writeSimpleElements(p *nexus.Printer, tags ...string) {
	for _, Name := range tags {
		funcName := capitalize(Name)
		p.Printf(
			`// %s returns a <%s/> element with optional attributes
func %s(c ...interface{}) *Element {
    return NewSimpleElement("%s", c...)
}`,
			funcName, string(Name), funcName, string(Name),
		)
		p.Println()
	}
}

func writeAttributes(p *nexus.Printer, names ...string) {
	for _, Name := range names {
		funcName := capitalize(Name)
		p.Printf(
			`// %s returns a %s="v" attribute
func %s(v string) *Attribute { return &Attribute{Name: %q, Val: v} }`,
			funcName, Name, funcName, Name,
		)
		p.Println()
	}
}

// ----------------------------------------

func capitalize(Name string) string {
	funcName := []byte(Name)
	funcName[0] = bytes.ToUpper(funcName)[0]
	return string(funcName)
}
