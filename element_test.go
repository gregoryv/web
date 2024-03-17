package web

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/gregoryv/asserter"
)

func Example_WalkElements() {
	root := Article(
		H1(),
		H2(),
		H3(),
	)
	WalkElements(root, func(e *Element) {
		fmt.Println(e.Name)
	})
	// output:
	// article
	// h1
	// h2
	// h3
}

func Test_elements(t *testing.T) {
	ok := func(el interface{}, exp ...string) {
		t.Helper()
		assert := asserter.New(t)
		w := bytes.NewBufferString("")
		hw := NewHtmlEncoder(w)
		switch el := el.(type) {
		case *Element:
			hw.Encode(el)
			got := el.String()
			for _, exp := range exp {
				assert().Contains(got, exp)
			}
		default:
			hw.Encode(el)
		}

		got := w.String()
		for _, exp := range exp {
			assert().Contains(got, exp)
		}
	}

	ok(Comment("hello", B("")), "<!--hello<b></b>-->")

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
	ok(Em(), "<em>", "</em>")
	ok(Fieldset(), "<fieldset>", "</fieldset>")
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
	ok(Input(), "<input/>")
	ok(Ins(), "<ins>", "</ins>")
	ok(Kbd(), "<kbd>", "</kbd>")
	ok(Label(), "<label>", "</label>")
	ok(Legend(), "<legend>", "</legend>")
	ok(Li(), "<li>", "</li>")
	ok(Main(), "<main>", "</main>")	
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

	// io.WriterTo
	ok(anyWriterTo("hello"), "hello")

	// Attributes
	ok(Action("x"), `action="x"`)
	ok(Attr("myown", "something"), `myown="something"`)
	ok(Attr("myown", 1), `myown="1"`)
	ok(Autocomplete("x"), `autocomplete="x"`)
	ok(Charset("x"), `charset="x"`)
	ok(Class("x"), `class="x"`)
	ok(Content("x"), `content="x"`)
	ok(For("x"), `for="x"`)	
	ok(Formaction("x"), `formaction="x"`)
	ok(Href("x"), `href="x"`)
	ok(Id("x"), `id="x"`)
	ok(Lang("x"), `lang="x"`)
	ok(Max("x"), `max="x"`)
	ok(Maxlength("x"), `maxlength="x"`)
	ok(Method("x"), `method="x"`)
	ok(Min("x"), `min="x"`)
	ok(Name("x"), `name="x"`)
	ok(OnBlur("main()"), `onBlur="main()"`)
	ok(OnFocus("main()"), `onFocus="main()"`)
	ok(OnLoad("main()"), `onLoad="main()"`)
	ok(OnMouseDown("main()"), `onMouseDown="main()"`)
	ok(OnMouseOut("main()"), `onMouseOut="main()"`)
	ok(OnMouseOver("main()"), `onMouseOver="main()"`)
	ok(OnMouseUp("main()"), `onMouseUp="main()"`)
	ok(OnMouseWheel("main()"), `onMouseWheel="main()"`)
	ok(Pattern("x"), `pattern="x"`)	
	ok(Placeholder("x"), `placeholder="x"`)
	ok(Rel("x"), `rel="x"`)
	ok(Size("x"), `size="x"`)
	ok(Src("x"), `src="x"`)
	ok(Tabindex("x"), `tabindex="x"`)
	ok(Type("x"), `type="x"`)
	ok(Value("x"), `value="x"`)
}

type anyWriterTo string

func (s anyWriterTo) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(s))
	return int64(n), err
}

func TestElement_Text(t *testing.T) {
	e := H2("my", Span("fancy"), "title")
	got := e.Text()
	exp := "my fancy title"
	assert := asserter.New(t)
	assert().Equals(got, exp)
}

func TestElement_HasAttr(t *testing.T) {
	e := Div(Class("x"), "text")
	if !e.HasAttr("class") {
		t.Error("attribute not found")
	}
	if e.HasAttr("other") {
		t.Error("found non existing attribute")
	}
}
