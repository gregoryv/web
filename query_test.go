package web

import (
	"fmt"
	"os"
	"testing"
)

func ExampleQuery() {
	doc := Article(
		H1("1"),
		H2("1.1"),
		H2("1.2"),
		H2("1.3", Class("mark")),
		H2("1.4"),
	)
	for _, el := range Query(doc, "h1") {
		el.With(Class("mark"))
	}

	for _, el := range Query(doc, "h2.mark") {
		el.WriteTo(os.Stdout)
	}

	fmt.Println()
	for _, el := range Query(doc, ".mark") {
		el.WriteTo(os.Stdout)
	}
	// output:
	// <h2 class="mark">1.3</h2>
	//
	// <h1 class="mark">1</h1>
	// <h2 class="mark">1.3</h2>
}

func TestQuery(t *testing.T) {
	doc := Article(
		H1("1"),
		H2("1.1"),
		H2("1.2", Id("x")),
		H2("1.3", Class("third")),
	)

	if got := Query(doc, "h2"); len(got) != 3 {
		t.Error(got)
	}

	if got := Query(doc, "h2.third"); len(got) != 1 {
		t.Error(got)
	}

	if got := Query(doc, "#x"); len(got) != 1 {
		t.Error(got)
	}
}
