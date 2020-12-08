package web

import "fmt"

type Hn struct {
	start int // root level used, if 1 then H1 == H1
}

func (me *Hn) H1(v ...interface{}) *Element { return me.hn(me.start, v...) }
func (me *Hn) H2(v ...interface{}) *Element { return me.hn(me.start+1, v...) }
func (me *Hn) H3(v ...interface{}) *Element { return me.hn(me.start+2, v...) }
func (me *Hn) H4(v ...interface{}) *Element { return me.hn(me.start+4, v...) }
func (me *Hn) H5(v ...interface{}) *Element { return me.hn(me.start+5, v...) }
func (me *Hn) H6(v ...interface{}) *Element { return me.hn(me.start+6, v...) }

func (me *Hn) hn(lev int, v ...interface{}) *Element {
	return NewElement(fmt.Sprintf("h%v", lev), v...)
}
