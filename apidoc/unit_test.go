package apidoc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gregoryv/must"
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

func Test_generate_apidoc(t *testing.T) {
	doc := NewDoc(NewRouter())

	nav := Nav()

	apis := Article(
		H1("API example documentation"),
		P("Plain and easy to read HTML documentation of your HTTP APIs"),
		"Table of contents",
		nav,

		H2("Path /"),
		P("Root resource users"),

		H3("List all users"),
		doc.NewRequest("GET", "/", nil),
		doc.JsonResponse(),

		H3("Filter user by name"),
		doc.NewRequest("GET", "/?name=John", nil),
		doc.Response(),

		doc.NewRequest("GET", "/?name=Whyat", nil),
		doc.Response(),

		P("It's also possible to filter using the POST method"),
		doc.NewRequest(
			"POST", "/", strings.NewReader(`{"name":"John"}`),
		),
		doc.JsonResponse(),

		H3("Accept text/html"),
		doc.Use(func() *http.Request {
			r := must.NewRequest("GET", "/?name=John", nil)
			r.Header.Set("Accept", "text/html")
			return r
		}()),
		doc.Response(),

		H2("Undefined path"),
		doc.NewRequest("GET", "/unknown/", nil),
		doc.Response(),

		// Section describing response headers and fields
		H2("Headers"),
		H3("Correlation-Id"),
		P(`Random string uniq for each request`),

		H2("Response"),
		H3("users"),
		P(`Array of user object, may be empty`),

		H4("name"),
		P(`String of max len 64`),

		H4("age"),
		P(`Uint8 `),
	)
	toc.MakeTOC(nav, apis, "h2", "h3")
	toc.GenerateIDs(apis, "h4")

	// Link words to the response and headers section
	refs := map[string]string{
		"users": "#users",
		"name":  "#name",
		"age":   "#age",
		// headers
		"Correlation-Id": "#correlationid",
	}
	LinkAll(apis, refs)

	page := NewFile("api_example.html",
		Html(
			Head(
				Meta(Charset("utf-8")),
				DefaultStyle(),
			),
			Body(
				apis,
			),
		),
	)
	page.SaveTo("../docs")
}

func NewRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/", ServeIndex)
	router.HandleFunc("/unknown/",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		},
	)
	return router
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	var m interface{}
	name := r.URL.Query().Get("name")
	if r.Method == "POST" {
		dec := json.NewDecoder(r.Body)
		var in struct {
			Name string `json:"name"`
		}
		dec.Decode(&in)
		name = in.Name
	}
	switch {
	case name == "John":
		m = struct {
			Users interface{} `json:"users"`
		}{data[0:1]}

	case name == "Whyat":
		m = struct {
			Users interface{} `json:"users"`
		}{data[1:]}

	default:
		m = struct {
			Users interface{} `json:"users"`
		}{data}
	}
	switch r.Header.Get("Accept") {
	case "text/html":
		w.Header().Set("Content-Type", "text/html")
		tmpl, _ := template.New("").Parse("{{.}}")
		var buf bytes.Buffer
		NewPage(Html(Body(fmt.Sprintf("%v", m)))).WriteTo(&buf)
		tmpl.Execute(w, buf.String())
	default:
		w.Header().Set("Content-Type", "application/json")
		// some random header
		cid := 6589408076449598657
		w.Header().Set("Correlation-Id", fmt.Sprintf("%d", cid))
		enc := json.NewEncoder(w)
		enc.Encode(&m)
	}
}

func init() { rand.Seed(time.Now().Unix()) }

var data = []struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}{
	{"John Doe", 42},
	{"Mat Whyat", 14},
	{"Lisa Whyat", 18},
}
