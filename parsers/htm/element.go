package htm

import (
	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

type Element = remark.Element
type Elements = remark.Elements

func NewElement(n *html.Node) (Element, error) {
	switch n.Data {
	case "a":
		return NewLink(n)
	case "code":
		return NewCode(n)
	case "em", "i":
		n.Data = "em"
		return NewGeneric(n)
	case "h1", "h2", "h3", "h4", "h5":
		return NewHeading(n)
	case "html":
		return NewHTML(n)
	case "img":
		return NewImage(n)
	case "ol":
		return NewOrderedList(n)
	case "strong", "b":
		n.Data = "strong"
		return NewGeneric(n)
	case "table":
		return NewTable(n)
	case "ul":
		return NewUnorderedList(n)
	}

	if n.Type == html.TextNode {
		return remark.String(n.Data), nil
	}

	return NewGeneric(n)
}
