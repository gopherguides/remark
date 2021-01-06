package htm

import (
	"fmt"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

type Element = remark.Element
type Elements = remark.Elements

func NewElement(n *html.Node) (Element, error) {
	switch n.Data {
	case "a":
		return NewLink(n)
	case "blockquote":
	case "body":
		return NewBody(n)
	case "code":
	case "div":
		return NewDiv(n)
	case "em", "i":
	case "h1", "h2", "h3", "h4", "h5":
	case "head":
		return NewHead(n)
	case "hr":
	case "html":
		return NewHTML(n)
	case "img":
		return NewImage(n)
	case "li":
	case "ol":
	case "p":
		return NewParagraph(n)
	case "row":
	case "strike":
	case "strong", "b":
	case "sub":
	case "sup":
	case "table":
	case "tbody":
	case "td":
	case "th":
	case "tr":
	case "ul":
	}

	if n.Type == html.TextNode {
		return remark.String(n.Data), nil
	}
	return nil, fmt.Errorf("unknown tag %s", n.Data)
}
