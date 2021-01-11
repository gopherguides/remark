package htm

import (
	"strings"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

type Element = remark.Element
type Elements = remark.Elements

func (p *Parser) NewElement(n *html.Node) (Element, error) {
	n.Data = strings.ToLower(n.Data)
	switch n.Data {
	case "hr":
		return &remark.HR{
			Attributes: remark.Attributes{},
		}, nil
	case "a":
		return p.NewLink(n)
	case "code":
		return p.NewCodeBlock(n)
	case "em", "i":
		n.Data = "em"
		return p.NewGeneric(n)
	case "h1", "h2", "h3", "h4", "h5":
		return p.NewHeading(n)
	case "img":
		return p.NewImage(n)
	case "ol":
		return p.NewOrderedList(n)
	case "strong", "b":
		n.Data = "strong"
		return p.NewGeneric(n)
	case "table":
		return p.NewTable(n)
	case "ul":
		return p.NewUnorderedList(n)
	}

	if n.Type == html.TextNode {
		s := strings.TrimSpace(n.Data)
		return remark.String(s), nil
	}

	return p.NewGeneric(n)
}
