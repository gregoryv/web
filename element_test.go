package web

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_elements(t *testing.T) {
	ok := func(el interface{}, exp ...string) {
		t.Helper()
		assert := asserter.New(t)
		w := bytes.NewBufferString("")
		hw := NewHtmlWriter(w)
		switch el := el.(type) {
		case *Element:
			hw.WriteHtml(el)
			got := el.String()
			for _, exp := range exp {
				assert().Contains(got, exp)
			}
		case *Attribute:
			hw.WriteHtml(el)
		default:
			t.Fatal(fmt.Errorf("unrecognized %#v", el))
		}

		got := w.String()
		for _, exp := range exp {
			assert().Contains(got, exp)
		}
	}

	ok(Html(), "<html>", "</html>")
	ok(Dl(Li("x")), "<dl>", "</dl>", "<li>", "</li>")
	ok(Link(Rel("x")), "link", `"x"`)
	ok(P("x", I("i")), "<p>", "x", "<i>")
	ok(A(), "<a>", "</a>")
	ok(Abbr(), "<abbr>", "</abbr>")
	ok(Acronym(), "<acronym>", "</acronym>")
	ok(Address(), "<address>", "</address>")
	ok(Article(), "<article>", "</article>")
	ok(Aside(), "<aside>", "</aside>")
	ok(B(), "<b>", "</b>")
	ok(Big(), "<big>", "</big>")
	ok(Blockquote(), "<blockquote>", "</blockquote>")
	ok(Body(), "<body>", "</body>")
	ok(Button(), "<button>", "</button>")
	ok(Cite(), "<cite>", "</cite>")
	ok(Code(), "<code>", "</code>")
	ok(Dd(), "<dd>", "</dd>")
	ok(Del(), "<del>", "</del>")
	ok(Details(), "<details>", "</details>")
	ok(Dfn(), "<dfn>", "</dfn>")
	ok(Div(), "<div>", "</div>")
	ok(Dl(), "<dl>", "</dl>")
	ok(Dt(), "<dt>", "</dt>")
	ok(Footer(), "<footer>", "</footer>")
	ok(Form(), "<form>", "</form>")
	ok(H1(), "<h1>", "</h1>")
	ok(H2(), "<h2>", "</h2>")
	ok(H3(), "<h3>", "</h3>")
	ok(H4(), "<h4>", "</h4>")
	ok(H5(), "<h5>", "</h5>")
	ok(H6(), "<h6>", "</h6>")
	ok(Head(), "<head>", "</head>")
	ok(Header(), "<header>", "</header>")
	ok(Hgroup(), "<hgroup>", "</hgroup>")
	ok(I(), "<i>", "</i>")
	ok(Ins(), "<ins>", "</ins>")
	ok(Kbd(), "<kbd>", "</kbd>")
	ok(Label(), "<label>", "</label>")
	ok(Legend(), "<legend>", "</legend>")
	ok(Li(), "<li>", "</li>")
	ok(Mark(), "<mark>", "</mark>")
	ok(Menu(), "<menu>", "</menu>")
	ok(Meter(), "<meter>", "</meter>")
	ok(Nav(), "<nav>", "</nav>")
	ok(Noscript(), "<noscript>", "</noscript>")
	ok(Ol(), "<ol>", "</ol>")
	ok(Optgroup(), "<optgroup>", "</optgroup>")
	ok(Option(), "<option>", "</option>")
	ok(Output(), "<output>", "</output>")
	ok(P(), "<p>", "</p>")
	ok(Pre(), "<pre>", "</pre>")
	ok(Quote(), "<quote>", "</quote>")
	ok(Script(), "<script>", "</script>")
	ok(Section(), "<section>", "</section>")
	ok(Select(), "<select>", "</select>")
	ok(Span(), "<span>", "</span>")
	ok(Style(), "<style>", "</style>")
	ok(Sub(), "<sub>", "</sub>")
	ok(Summary(), "<summary>", "</summary>")
	ok(Sup(), "<sup>", "</sup>")
	ok(Table(), "<table>", "</table>")
	ok(Tbody(), "<tbody>", "</tbody>")
	ok(Td(), "<td>", "</td>")
	ok(Textarea(), "<textarea>", "</textarea>")
	ok(Th(), "<th>", "</th>")
	ok(Thead(), "<thead>", "</thead>")
	ok(Title(), "<title>", "</title>")
	ok(Tr(), "<tr>", "</tr>")
	ok(U(), "<u>", "</u>")
	ok(Ul(), "<ul>", "</ul>")
	ok(Var(), "<var>", "</var>")
	ok(Base(), "<base/>")
	ok(Br(), "<br/>")
	ok(Hr(), "<hr/>")
	ok(Img(), "<img/>")
	ok(Input(), "<input/>")
	ok(Keygen(), "<keygen/>")
	ok(Link(), "<link/>")
	ok(Meta(), "<meta/>")
	// Attributes
	ok(Charset("x"), `charset="x"`)
	ok(Class("x"), `class="x"`)
	ok(Content("x"), `content="x"`)
	ok(Href("x"), `href="x"`)
	ok(Id("x"), `id="x"`)
	ok(Lang("x"), `lang="x"`)
	ok(Name("x"), `name="x"`)
	ok(Rel("x"), `rel="x"`)
	ok(Src("x"), `src="x"`)
	ok(Type("x"), `type="x"`)
}