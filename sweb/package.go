/*
Package sweb provides sequential style of writing web elements.

The package is not safe to use concurrently as there is a global stack
handling the state.

*/
package sweb

import (
	"github.com/gregoryv/web"
)

var stack []*web.Element

func push(v *web.Element) {
	stack = append(stack, v)
}

func pop() *web.Element {
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return v
}

func parent() *web.Element {
	if len(stack) == 0 {
		return nil
	}
	return stack[len(stack)-1]
}

// end pops the stack until the named element is popped
func end(name string) *web.Element {
	el := pop()
	if el.Name != name {
		return end(name)
	}
	return el
}

// End unwinds the stack returning the first element
func End() *web.Element {
	if len(stack) == 0 {
		return nil
	}
	v := stack[0]
	stack = stack[:0] // clear
	return v
}
