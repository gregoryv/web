package apidoc_test

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gregoryv/web/apidoc"
)

func ExampleNewIntercepter() {
	mux := http.NewServeMux()
	x := apidoc.NewIntercepter(mux)

	// use the intercepter when defining routes
	x.Handle("/", someHandler)
	x.Handle("POST /{id}", someHandler)

	fmt.Println(x.Routes())
	// output:
	// [GET / POST /{id}]
}

func ExampleIntercepter_Defines() {
	mux := http.NewServeMux()
	x := apidoc.NewIntercepter(mux)
	x.Handle("/", someHandler)
	x.Handle("POST /", someHandler)

	// Optionally set the ErrHandler, e.g. in your test
	// x.ErrHandler = t

	// Call Defines when you document the routes
	x.Defines("GET /")
	x.Defines("GET /nosuch/thing")

	// Check if all routes have been documented
	fmt.Println(x.Undocumented())
	// output:
	// Defines("GET /nosuch/thing"): no such route
	// POST /
}

func ExampleDoc_JsonResponse() {
	doc := apidoc.NewDoc(http.HandlerFunc(someRouter))

	doc.NewRequest("GET", "/", nil).WriteTo(os.Stdout)
	doc.JsonResponse().WriteTo(os.Stdout)
	// output:
	// <pre class="request">HTTP/1.1 GET /
	// </pre><pre class="response">HTTP/1.1 200 OK
	//
	// {
	//     "animal": "Goat",
	//     "age": 10,
	//     "friendly": "hell no, not this one"
	// }</pre>
}

func ExampleJsonResponseFrom() {
	r, _ := http.NewRequest("GET", "/", nil)

	apidoc.JsonResponseFrom(
		http.HandlerFunc(someRouter),
		r,
	).WriteTo(os.Stdout)
	// output:
	// <pre class="response">HTTP/1.1 200 OK
	//
	// {
	//     "animal": "Goat",
	//     "age": 10,
	//     "friendly": "hell no, not this one"
	// }</pre>
}

func ExampleRawResponseFrom() {
	r, _ := http.NewRequest("GET", "/", nil)
	element := apidoc.RawResponseFrom(
		http.HandlerFunc(someRouter),
		r,
	)
	element.WriteTo(os.Stdout)
	// output:
	// <pre class="response">HTTP/1.1 200 OK
	//
	// {"animal": "Goat","age": 10, "friendly": "hell no, not this one"}</pre>
}

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

var someHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
})
