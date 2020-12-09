package web

// Wrap returns an element that produces no output when rendered apart
// from it's children.
func Wrap(v ...interface{}) *Element {
	return NewElement("wrapper", v...)
}
