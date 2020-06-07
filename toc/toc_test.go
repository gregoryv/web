package toc

import (
	"os"

	. "github.com/gregoryv/web"
)

func Example() {
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
