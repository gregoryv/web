package web

import (
	"regexp"
	"strings"
)

// ILinkAll replaces key words same as LinkAll but is canse insensitive
func ILinkAll(root *Element, refs map[string]string) {
	linkAll(root, refs, false)
}

// LinkAll replaces key words found in the root and it's
// children with links defined in the map. The map should be TEXT -> HREF
func LinkAll(root *Element, refs map[string]string) {
	linkAll(root, refs, true)
}

// LinkAll replaces key words found in the root and it's
// children with links defined in the map. The map should be TEXT -> HREF
func linkAll(root *Element, refs map[string]string, caseSensitive bool) {
	WalkElements(root, func(e *Element) {
		if e.Name == "a" { // inside an link already
			return
		}
		for i, c := range e.Children {
			if c, ok := c.(string); ok {
				for txt, href := range refs {
					txt := strings.ReplaceAll(txt, " ", `[\n\t\r\s]*`)
					exp := `\b(` + txt + `)\b`
					if !caseSensitive {
						exp = "(?i)" + exp
					}
					re := regexp.MustCompile(exp)
					c = re.ReplaceAllString(c, `<a href="`+href+`">$1</a>`)
				}
				e.Children[i] = c
			}
		}
	})
}
