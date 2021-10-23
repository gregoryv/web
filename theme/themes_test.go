package theme

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/gregoryv/english"
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

var variants = []struct {
	name string
	fn   func() *CSS
}{
	{
		"GoldenSpace",
		GoldenSpace,
	},
	{
		"GoldenSpace_GoishColors", func() *CSS {
			return GoldenSpace().With(GoishColors())
		},
	},
}

func Test_generate_pages(t *testing.T) {
	rand.Seed(0)

	body := testBody()
	for _, c := range variants {
		t.Run(c.name, func(t *testing.T) {
			page := NewSafePage(
				Html(
					Style(c.fn()),
					body,
				),
			)
			out := fmt.Sprintf("examples/%s.html", c.name)
			page.SaveAs(out)
			t.Log("generated theme example: ", out)
		})
	}
}

func testBody() *Element {
	article := Article(
		Section(
			H2(title()),
			P(randomSentences(3), "p"),
			Ol(
				Li("red"),
				Li("blue"),
				Li("green"),
			),

			H2(title()),
			P(randomSentences(5)),

			H3(title()),
			P(randomSentences(5)),
		),
		Section(
			H2(title()),

			P(
				randomSentences(1), " ",
				Code(`var x string = "..."`), " ",
				randomSentences(1),
			),

			Pre(`package yours

import (
    . "github.com/gregoryv/web"
)
// ...`),

			Table(
				Thead(
					Tr(
						Th("Nouns"),
						Th("Adjectives"),
					),
				),
				Tbody(
					func() *Element {
						v := Wrap()
						for i := 0; i < 10; i++ {
							v.With(
								Tr(
									Td(english.ClassNoun.Random()),
									Td(english.ClassAdjective.Random()),
								),
							)
						}
						return v
					}(),
				),
			),
		),
	)

	nav := Nav()
	toc.MakeTOC(nav, article, "h2", "h3")

	body := Body(
		Header(
			"Header: Theme examples by the gregoryv/web/theme package",

			func() *Element {
				ul := Ul()
				for _, c := range variants {
					ul.With(
						Li(A(Href(c.name+".html"), c.name)),
					)
				}
				return ul
			}(),

			H1(title()),
			P(randomSentences(7)),
		),
		"Table of contents",
		Br(),
		nav,
		article,
		Hr(),
		Footer(
			"Footer: Author Gregory Vincic",
		),
	)
	return body
}

func randomWords(n int) string {
	words := english.RandomWords(n)
	return strings.Join(words, " ")
}

func title() string {
	s := english.RandomStatement(3, 5)
	return english.Sentence(s, ' ')
}

func sentence() string {
	return randomSentences(1)
}

func randomSentences(n int) string {
	sentences := make([]string, n)
	for i, _ := range sentences {
		s := english.RandomStatement(4, 8)
		sentences[i] = english.Sentence(s, '.')
	}

	return strings.Join(sentences, " ")
}
