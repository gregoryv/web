package toc

import (
	"os"

	. "github.com/gregoryv/web"
)

func ExampleMakeTOC() {
	nav := Nav()
	a := Article(
		H1("Programming"),
		nav,
		H2("Design"),
		H3("Diagrams"),

		Section(
			H2(Id("myid"), "Test"),
			H3("Unit"),
			H3("Integration"),
		),
	)
	MakeTOC(nav, a, "h2")
	nav.WriteTo(os.Stdout)
	// output:
	// <nav><ul>
	// <li class="h2"><a href="#design">Design</a></li>
	// <li class="h2"><a href="#myid">Test</a></li>
	// </ul>
	// </nav>
}

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
	// <li class="h2"><a href="#design">Design</a></li>
	// <li class="h2"><a href="#myid">Test</a></li>
	// </ul>
}

func ExampleGenerateIDs() {
	a := Article(
		Section(
			H2(Id("current"), "Current car"),
			H2("My first car"),
			H3("Broke down"),
		),
	)
	GenerateIDs(a, "h2")
	a.WriteTo(os.Stdout)
	// output:
	// <article>
	// <section>
	// <h2 id="current">Current car</h2>
	// <h2 id="myfirstcar">My first car</h2>
	// <h3>Broke down</h3>
	// </section>
	// </article>
}
