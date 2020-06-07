package toc

import (
	"os"

	. "github.com/gregoryv/web"
)

func ExampleParseTOC() {
	a := Article(
		H1("Programming"),
		H2("Design"),
		H3("Diagrams"),

		Section(
			H2(Id("myid"), "Test"),
			H3("Unit"),
			H3("Integration"),
		),
	)
	toc := ParseTOC(a, "h2")
	toc.WriteTo(os.Stdout)
	// output:
	// <ul>
	// <li><a href="#design">Design</a></li>
	// <li><a href="#myid">Test</a></li>
	// </ul>
}

func ExampleGenerateIds() {
	a := Article(
		H2("My first car"),
		H3("Broke down"),
	)
	GenerateIds(a, "h2")
	a.WriteTo(os.Stdout)
	// output:
	// <article>
	// <h2 id="myfirstcar">My first car</h2>
	// <h3>Broke down</h3>
	// </article>
}
