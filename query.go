package web

import "strings"

// Query returns elements matching the given css selector style
// expression. See ParseExpr for supported expressions.
func Query(root *Element, expr string) []*Element {
	return Find(root, ParseExpr(expr)...)
}

// Find returns the matching elements of the expression
func Find(root *Element, matchers ...Matcher) []*Element {
	res := make([]*Element, 0)

	WalkElements(root, func(el *Element) {
		for _, match := range matchers {
			if !match(el) {
				return
			}
		}
		res = append(res, el)
	})
	return res
}

// ParseExpr returns a list of matchers to use in func Find.
//
// Valid expressions:
//   Name
//   Name.Class
//   .Class
func ParseExpr(expr string) []Matcher {
	res := make([]Matcher, 0)

	parts := strings.Split(expr, ".")

	switch {
	case len(parts) == 1: // name only
		res = append(res, func(el *Element) bool {
			return el.Name == expr
		})
	case len(parts) == 2 && parts[0] != "": // name and class
		res = append(res, func(el *Element) bool {
			name := parts[0]
			class := parts[1]
			return el.Name == name && el.AttrVal("class") == class
		})
	case len(parts) == 2 && parts[0] == "": // class
		res = append(res, func(el *Element) bool {
			return el.AttrVal("class") == parts[1]
		})
	}

	return res
}

type Matcher func(*Element) bool
