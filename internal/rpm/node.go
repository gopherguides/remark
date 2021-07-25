package rpm

import "golang.org/x/net/html"

func NewHTMLNode(name string, attrs Attrs) *html.Node {
	n := &html.Node{
		Data: name,
		Attr: attrs.ToAttrs(),
	}
	return n
}
