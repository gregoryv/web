package nexus

import (
	"fmt"
	"io"
)

// NewPrinter returns a writer nexus. The referenced error is set by
// all methods if an error occurs.
func NewPrinter(w io.Writer) (*Printer, *error) {
	t := &Printer{w: w}
	return t, &t.err
}

type Printer struct {
	w   io.Writer
	err error
	// Total number of bytes written
	Written int64
}

// Print prints arguments using the underlying writer. Does nothing if
// Printer has failed.
func (t *Printer) Print(args ...interface{}) {
	if t.err != nil {
		return
	}
	var n int
	n, t.err = fmt.Fprint(t.w, args...)
	t.Written += int64(n)
}

// Printf prints a formated string using the underlying writer. Does
// nothing if Printer has failed.
func (t *Printer) Printf(format string, args ...interface{}) {
	if t.err != nil {
		return
	}
	var n int
	n, t.err = fmt.Fprintf(t.w, format, args...)
	t.Written += int64(n)
}

// Println prints arguments using the underlying writer. Does nothing if
// Printer has failed.
func (t *Printer) Println(args ...interface{}) {
	if t.err != nil {
		return
	}
	var n int
	n, t.err = fmt.Fprintln(t.w, args...)
	t.Written += int64(n)
}

// Write writes the bytes using the underlying writer. Does nothing if
// Printer has failed.
func (t *Printer) Write(b []byte) (int, error) {
	if t.err != nil {
		return 0, t.err
	}
	var n int
	n, t.err = t.w.Write(b)
	t.Written += int64(n)
	return n, t.err
}
