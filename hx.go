package web

import "fmt"

// NewHn returns a Hn starting at the given level. 1 meaning H1 produces H1, whereas
// 2 means H1 produces H2 and so on.
func NewHn(start int) *Hn {
	if start <= 0 {
		start = 1
	}
	return &Hn{start - 1}
}

// Hn allows for creating headings starting at different levels.  Zero
// value object of Hn is usable, where H1 method matches H1 heading.
type Hn struct {
	start int // root level used, if 0 then H1 == H1
}

func (me *Hn) H1(v ...interface{}) *Element { return me.hn(1, v...) }
func (me *Hn) H2(v ...interface{}) *Element { return me.hn(2, v...) }
func (me *Hn) H3(v ...interface{}) *Element { return me.hn(3, v...) }
func (me *Hn) H4(v ...interface{}) *Element { return me.hn(4, v...) }
func (me *Hn) H5(v ...interface{}) *Element { return me.hn(5, v...) }
func (me *Hn) H6(v ...interface{}) *Element { return me.hn(6, v...) }

func (me *Hn) hn(lev int, v ...interface{}) *Element {
	return NewElement(fmt.Sprintf("h%v", me.start+lev), v...)
}
