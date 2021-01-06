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
		return NewBlockQuote(n)
	case "body":
		return NewBody(n)
	case "code":
		return NewCode(n)
	case "div":
		return NewDiv(n)
	case "em", "i":
		return NewEm(n)
	case "h1", "h2", "h3", "h4", "h5":
		return NewHeading(n)
	case "head":
		return NewHead(n)
	case "hr":
		return NewHR(n)
	case "html":
		return NewHTML(n)
	case "img":
		return NewImage(n)
	case "li":
		return NewListItem(n)
	case "ol":
		return NewOrderedList(n)
	case "p":
		return NewParagraph(n)
	case "row":
		return NewRow(n)
	case "strike":
		return NewStrike(n)
	case "strong", "b":
		return NewStrong(n)
	case "sub":
		return NewSub(n)
	case "sup":
		return NewSup(n)
	case "table":
		return NewTable(n)
	case "tbody":
		return NewTableBody(n)
	case "td":
		return NewTD(n)
	case "th":
		return NewTH(n)
	case "tr":
		return NewTR(n)
	case "ul":
		return NewUnorderedList(n)
	}

	if n.Type == html.TextNode {
		return remark.String(n.Data), nil
	}
	return nil, fmt.Errorf("unknown tag %s", n.Data)
}
