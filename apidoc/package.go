/*
   Package apidoc is a way to define self documenting HTTP API's.

*/
package apidoc

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"text/template"

	"github.com/gregoryv/ex"
	"github.com/gregoryv/must"
	. "github.com/gregoryv/web"
)

// NewDoc returns a documentation generator for your router
func NewDoc(router http.Handler) *Doc {
	return &Doc{
		router: router,
	}
}

type Doc struct {
	router http.Handler
	*http.Request
}

var DefaultStyle = Style(`
html, body { margin: 0 0; padding: 0 0; }
body {
  padding: 1em 1.618em 1em 1.618em;
}
h1:first-child {
  margin-top: 0;
}
.request {
  padding: 1em 1.618em;
}

.response {
  padding: 1em 1.618em;
  background-color: #f2f2f2;
  border-radius: 1em;
}
`)

func (d *Doc) NewRequest(method, path string, body io.Reader) *Element {
	r := must.NewRequest(method, path, body)
	d.Request = r
	return RawRequest(r)
}

func (d *Doc) Response() *Element {
	return RawResponseFrom(d.router, d.Request)
}

func (d *Doc) JsonResponse() *Element {
	return JsonResponseFrom(d.router, d.Request)
}

func (d *Doc) Use(r *http.Request) *Element {
	d.Request = r
	return RawRequest(r)
}

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

func RawResponseFrom(h http.Handler, r *http.Request) *Element {
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	resp := w.Result()
	return RawResponse(resp)
}

func RawResponse(resp *http.Response) *Element {
	var (
		data, _ = ioutil.ReadAll(resp.Body)
		body    bytes.Buffer
		headers bytes.Buffer
	)
	template.HTMLEscape(&body, data)

	resp.Header.Write(&headers)
	return Pre(Class("response"),
		statusLine(resp),
		headers.String(),
		"\n",
		body.String(),
	)
}

func JsonResponseFrom(h http.Handler, r *http.Request) *Element {
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	resp := w.Result()
	return JsonResponse(resp)
}

func JsonResponse(resp *http.Response) *Element {
	var (
		data, _ = ioutil.ReadAll(resp.Body)
		headers bytes.Buffer
		body    bytes.Buffer
	)
	resp.Header.Write(&headers)
	jw := ex.NewJsonWriter()
	jw.WriteTo(&body, []byte(data))
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
