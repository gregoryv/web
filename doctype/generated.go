package doctype

// generated by internal/dog.go
func A(c ...interface{}) *Tag {
	return NewTag("a", c...)
}
func Abbr(c ...interface{}) *Tag {
	return NewTag("abbr", c...)
}
func Acronym(c ...interface{}) *Tag {
	return NewTag("acronym", c...)
}
func Address(c ...interface{}) *Tag {
	return NewTag("address", c...)
}
func Article(c ...interface{}) *Tag {
	return NewTag("article", c...)
}
func Aside(c ...interface{}) *Tag {
	return NewTag("aside", c...)
}
func B(c ...interface{}) *Tag {
	return NewTag("b", c...)
}
func Big(c ...interface{}) *Tag {
	return NewTag("big", c...)
}
func Blockquote(c ...interface{}) *Tag {
	return NewTag("blockquote", c...)
}
func Body(c ...interface{}) *Tag {
	return NewTag("body", c...)
}
func Button(c ...interface{}) *Tag {
	return NewTag("button", c...)
}
func Cite(c ...interface{}) *Tag {
	return NewTag("cite", c...)
}
func Code(c ...interface{}) *Tag {
	return NewTag("code", c...)
}
func Dd(c ...interface{}) *Tag {
	return NewTag("dd", c...)
}
func Del(c ...interface{}) *Tag {
	return NewTag("del", c...)
}
func Details(c ...interface{}) *Tag {
	return NewTag("details", c...)
}
func Dfn(c ...interface{}) *Tag {
	return NewTag("dfn", c...)
}
func Div(c ...interface{}) *Tag {
	return NewTag("div", c...)
}
func Dl(c ...interface{}) *Tag {
	return NewTag("dl", c...)
}
func Dt(c ...interface{}) *Tag {
	return NewTag("dt", c...)
}
func Footer(c ...interface{}) *Tag {
	return NewTag("footer", c...)
}
func Form(c ...interface{}) *Tag {
	return NewTag("form", c...)
}
func H1(c ...interface{}) *Tag {
	return NewTag("h1", c...)
}
func H2(c ...interface{}) *Tag {
	return NewTag("h2", c...)
}
func H3(c ...interface{}) *Tag {
	return NewTag("h3", c...)
}
func H4(c ...interface{}) *Tag {
	return NewTag("h4", c...)
}
func H5(c ...interface{}) *Tag {
	return NewTag("h5", c...)
}
func H6(c ...interface{}) *Tag {
	return NewTag("h6", c...)
}
func Head(c ...interface{}) *Tag {
	return NewTag("head", c...)
}
func Header(c ...interface{}) *Tag {
	return NewTag("header", c...)
}
func Hgroup(c ...interface{}) *Tag {
	return NewTag("hgroup", c...)
}
func I(c ...interface{}) *Tag {
	return NewTag("i", c...)
}
func Ins(c ...interface{}) *Tag {
	return NewTag("ins", c...)
}
func Kbd(c ...interface{}) *Tag {
	return NewTag("kbd", c...)
}
func Label(c ...interface{}) *Tag {
	return NewTag("label", c...)
}
func Legend(c ...interface{}) *Tag {
	return NewTag("legend", c...)
}
func Li(c ...interface{}) *Tag {
	return NewTag("li", c...)
}
func Mark(c ...interface{}) *Tag {
	return NewTag("mark", c...)
}
func Menu(c ...interface{}) *Tag {
	return NewTag("menu", c...)
}
func Meter(c ...interface{}) *Tag {
	return NewTag("meter", c...)
}
func Nav(c ...interface{}) *Tag {
	return NewTag("nav", c...)
}
func Noscript(c ...interface{}) *Tag {
	return NewTag("noscript", c...)
}
func Ol(c ...interface{}) *Tag {
	return NewTag("ol", c...)
}
func Optgroup(c ...interface{}) *Tag {
	return NewTag("optgroup", c...)
}
func Option(c ...interface{}) *Tag {
	return NewTag("option", c...)
}
func Output(c ...interface{}) *Tag {
	return NewTag("output", c...)
}
func P(c ...interface{}) *Tag {
	return NewTag("p", c...)
}
func Pre(c ...interface{}) *Tag {
	return NewTag("pre", c...)
}
func Script(c ...interface{}) *Tag {
	return NewTag("script", c...)
}
func Section(c ...interface{}) *Tag {
	return NewTag("section", c...)
}
func Select(c ...interface{}) *Tag {
	return NewTag("select", c...)
}
func Span(c ...interface{}) *Tag {
	return NewTag("span", c...)
}
func Style(c ...interface{}) *Tag {
	return NewTag("style", c...)
}
func Sub(c ...interface{}) *Tag {
	return NewTag("sub", c...)
}
func Summary(c ...interface{}) *Tag {
	return NewTag("summary", c...)
}
func Sup(c ...interface{}) *Tag {
	return NewTag("sup", c...)
}
func Table(c ...interface{}) *Tag {
	return NewTag("table", c...)
}
func Tbody(c ...interface{}) *Tag {
	return NewTag("tbody", c...)
}
func Td(c ...interface{}) *Tag {
	return NewTag("td", c...)
}
func Textarea(c ...interface{}) *Tag {
	return NewTag("textarea", c...)
}
func Th(c ...interface{}) *Tag {
	return NewTag("th", c...)
}
func Thead(c ...interface{}) *Tag {
	return NewTag("thead", c...)
}
func Title(c ...interface{}) *Tag {
	return NewTag("title", c...)
}
func Tr(c ...interface{}) *Tag {
	return NewTag("tr", c...)
}
func U(c ...interface{}) *Tag {
	return NewTag("u", c...)
}
func Ul(c ...interface{}) *Tag {
	return NewTag("ul", c...)
}
func Var(c ...interface{}) *Tag {
	return NewTag("var", c...)
}

func Base(c ...interface{}) *Tag {
	return NewSimpleTag("base", c...)
}
func Br(c ...interface{}) *Tag {
	return NewSimpleTag("br", c...)
}
func Hr(c ...interface{}) *Tag {
	return NewSimpleTag("hr", c...)
}
func Img(c ...interface{}) *Tag {
	return NewSimpleTag("img", c...)
}
func Input(c ...interface{}) *Tag {
	return NewSimpleTag("input", c...)
}
func Keygen(c ...interface{}) *Tag {
	return NewSimpleTag("keygen", c...)
}
func Link(c ...interface{}) *Tag {
	return NewSimpleTag("link", c...)
}
func Meta(c ...interface{}) *Tag {
	return NewSimpleTag("meta", c...)
}

func Charset(v string) *Attr { return &Attr{name: "charset", val: v} }
func Content(v string) *Attr { return &Attr{name: "content", val: v} }
func Href(v string) *Attr    { return &Attr{name: "href", val: v} }
func Lang(v string) *Attr    { return &Attr{name: "lang", val: v} }
func Name(v string) *Attr    { return &Attr{name: "name", val: v} }
func Rel(v string) *Attr     { return &Attr{name: "rel", val: v} }
func Src(v string) *Attr     { return &Attr{name: "src", val: v} }
func Type(v string) *Attr    { return &Attr{name: "type", val: v} }
