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

func Test_generate_pages(t *testing.T) {
	rand.Seed(0)
	cases := map[string]func() *CSS{
		"GoldenSpace": GoldenSpace,
		"GoldenSpace_GoishColors": func() *CSS {
			return GoldenSpace().With(GoishColors())
		},
	}
	body := testBody()
	for name, theme := range cases {
		t.Run(name, func(t *testing.T) {
			page := NewSafePage(
				Html(
					Style(theme()),
					body,
				),
			)
			filename := fmt.Sprintf("examples/%s.html", name)
			page.SaveAs(filename)
			t.Log("generated theme example: ", filename)
		})
	}
}

func testBody() *Element {
	article := Article(
		Section(
			H2(title()),
			P(randomSentences(3), "p"),
			Ol(
				Li("Li: red"),
				Li("Li: blue"),
				Li("Li: green"),
			),

			H2(title()),
			P(randomSentences(5)),

			H3(title()),
			P(randomSentences(5)),
		),
	)

	nav := Nav()
	toc.MakeTOC(nav, article, "h2", "h3")

	body := Body(
		Header(
			"Header: Author Gregory Vincic",

			H1(title()),
			P(randomSentences(7)),
		),
		"Table of contents",
		Br(),
		nav,
		article,
		Footer(
			"Footer: ...",
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
