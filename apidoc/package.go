/*
   Package apidoc provides html document builder for http requests and responses.

*/
package apidoc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gregoryv/must"
	. "github.com/gregoryv/web"
)

// NewDoc returns a documentation generator for the given router
func NewDoc(router http.Handler) *Doc {
	return &Doc{
		router: router,
	}
}

type Doc struct {
	router http.Handler
	*http.Request
}

// NewRequest returns a <pre> element of a request based on the
// arguments. For more advanced requests use Doc.Use()
func (d *Doc) NewRequest(method, path string, body io.Reader) *Element {
	r := must.NewRequest(method, path, body)
	return d.Use(r)
}

// Response returns a raw response from the last used request.
func (d *Doc) Response() *Element {
	return RawResponseFrom(d.router, d.Request)
}

// JsonResponse returns a tidy json response from the last used request
func (d *Doc) JsonResponse() *Element {
	return JsonResponseFrom(d.router, d.Request)
}

// Use returns a <pre> element of the given request.
func (d *Doc) Use(r *http.Request) *Element {
	d.Request = r
	return RawRequest(r)
}

// RawRequest returns a <pre> element with the request. The request is
// reusable afterwards.
func RawRequest(r *http.Request) *Element {
	var headers bytes.Buffer
	r.Header.Write(&headers)
	pre := Pre(Class("request"),
		requestLine(r),
		headers.String(),
	)
	if r.Body != nil {
		pre.With(
			"\n",
			string(readRestoreBody(r)),
		)
	}
	return pre
}

func readRestoreBody(r *http.Request) []byte {
	var body []byte
	if r.Body != nil {
		body, _ = ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewReader(body))
	}
	return body
}

// RawResponseFrom returns the full response from the request to the given handler
func RawResponseFrom(h http.Handler, r *http.Request) *Element {
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	resp := w.Result()
	return RawResponse(resp)
}

// RawResponse dumps the response including body
func RawResponse(resp *http.Response) *Element {
	var (
		data, _ = ioutil.ReadAll(resp.Body)
		body    bytes.Buffer
		headers bytes.Buffer
	)
	body.Write(data)

	resp.Header.Write(&headers)
	return Pre(Class("response"),
		statusLine(resp),
		headers.String(),
		"\n",
		body.String(),
	)
}

// JsonResponseFrom records the response of the request on the handler
// and returns same as JsonResponse.
func JsonResponseFrom(h http.Handler, r *http.Request) *Element {
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	resp := w.Result()
	return JsonResponse(resp)
}

// JsonResponse converts the response to a <pre> element including the body.
func JsonResponse(resp *http.Response) *Element {
	var (
		data, _ = ioutil.ReadAll(resp.Body)
		headers bytes.Buffer
		body    bytes.Buffer
	)
	resp.Header.Write(&headers)
	json.Indent(&body, data, "", "    ")
	return Pre(Class("response"),
		statusLine(resp),
		headers.String(),
		"\n",
		body.String(),
	)
}

func requestLine(r *http.Request) string {
	return fmt.Sprintf("%s %s %s\n", r.Proto, r.Method, r.URL)
}

func statusLine(resp *http.Response) string {
	return fmt.Sprintf("%s %s\n", resp.Proto, resp.Status)
}

func DefaultStyle() *Element {
	css := NewCSS()
	css.Style("html, body",
		"margin: 0 0",
		"padding: 0 0",
	)
	css.Style("body",
		"padding: 1em 1.618em 1em 1.618em",
	)
	css.Style("h1:first-child",
		"margin-top: 0",
	)
	css.Style(".request",
		"padding: 1em 1.618em",
		"border-radius: 1em",
		"border: 1px dashed #929292",
	)
	css.Style(".response",
		"padding: 1em 1.618em",
		"background-color: #f2f2f2",
		"border-radius: 1em",
	)
	css.Style("nav ul",
		"list-style-type: none",
		"padding-left: 0",
	)
	css.Style("nav ul .h3",
		"margin-left: 1em",
	)
	return Style(css)
}
