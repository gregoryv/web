package apidoc_test

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gregoryv/web"
	"github.com/gregoryv/web/apidoc"
)

func someRouter(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"animal": "Goat","age": 10, "friendly": "hell no, not this one"}`)
	}
	if r.Method == "POST" {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"message": "added"}`)
	}
}

func ExampleNewDoc() {
	doc := apidoc.NewDoc(http.HandlerFunc(someRouter))

	body := web.Body(
		doc.NewRequest("GET", "/", nil),
		doc.JsonResponse(),
	)
	out := web.NewHtmlWriter(os.Stdout)
	out.WriteHtml(body)
	// output:
	// <body>
	// <pre class="request">HTTP/1.1 GET /
	// </pre><pre class="response">HTTP/1.1 200 OK
	//
	// {
	//     "animal": "Goat",
	//     "age": 10,
	//     "friendly": "hell no, not this one"
	// }</pre></body>

}
