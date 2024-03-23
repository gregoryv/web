package apidoc

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

// NewIntercepter returns a wrapper for the underlying handler
// registering all routes being defined. It can then be used to add
// more documentation to each route.
func NewIntercepter(impl Handler) *Intercepter {
	return &Intercepter{
		impl:       impl,
		index:      make(map[string]struct{}),
		ErrHandler: &logIt{},
	}
}

type Intercepter struct {
	impl   Handler
	routes []string
	index  map[string]struct{}

	// handles errors when calling method Defines
	ErrHandler
}

func (d *Intercepter) Handle(pattern string, h http.Handler) {
	key := WithMethod(pattern)
	d.routes = append(d.routes, key)
	d.index[key] = struct{}{}
	d.impl.Handle(pattern, h)
}

// Routes returns a list of defined routes as "METHOD PATTERN"
func (d *Intercepter) Routes() []string {
	return d.routes
}

// Defines checks if the given pattern, [METHOD ]PATTERN, has not been
// defined. Use it when documenting your routes.  The given error
// handler is used to signal error, eg. using testing.T in a test.
func (d *Intercepter) Defines(pattern string) {
	key := WithMethod(pattern)
	// if the ErrHandler is e.g. testing.T
	if eh, ok := d.ErrHandler.(interface{ Helper() }); ok {
		eh.Helper()
	}
	if _, found := d.index[key]; !found {
		d.ErrHandler.Errorf("Defines(\"%s\"): no such route", key)
		return
	}
	delete(d.index, key)
}

// Undocumented returns empty string if all routes are documented.
func (d *Intercepter) Undocumented() string {
	if len(d.index) == 0 {
		return ""
	}
	var buf bytes.Buffer
	for key, _ := range d.index {
		fmt.Fprintln(&buf, key)
	}
	return buf.String()
}

// WithMethod prefixes pattern with GET if it starts with /
func WithMethod(pattern string) string {
	if strings.HasPrefix(pattern, "/") {
		pattern = fmt.Sprintf("GET %s", pattern)
	}
	return pattern
}

// ----------------------------------------

type Router interface {
	Handler
	http.Handler
}

type Handler interface {
	Handle(pattern string, serve http.Handler)
}

// ----------------------------------------

type ErrHandler interface {
	Errorf(format string, args ...any)
}

type logIt struct{}

func (_ *logIt) Errorf(format string, args ...any) {
	fmt.Println(fmt.Sprintf(format, args...))
}
