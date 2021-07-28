package htm

import (
	"fmt"
	"strings"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func HTMLBody(root *html.Node) *html.Node {
	if root == nil {
		return nil
	}

	switch root.Data {
	case "body":
		return root
	case "html":
		return HTMLBody(root.FirstChild)
	case "head":
		return HTMLBody(root.NextSibling)
	case "":
		return HTMLBody(root.FirstChild)
	}
	return root

}

func (p *Parser) NewTag(n *html.Node) (remark.Tag, error) {
	if n == nil {
		return nil, fmt.Errorf("nil node")
	}

	n.Data = strings.TrimSpace(n.Data)
	if n.Type == html.TextNode {
		c := strings.TrimSpace(n.Data)
		if len(c) == 0 {
			return nil, nil
		}
		return remark.String(c), nil
	}

	n.Data = strings.ToLower(n.Data)
	if len(n.Data) == 0 {
		g, err := p.NewGeneric(n)
		if err != nil {
			return nil, err
		}
		return &Document{
			Generic: g,
		}, nil
	}

	switch n.Data {
	case "img":
		return p.NewImage(n)
	case "code":
		return p.NewCodeBlock(n)
	case "em", "i":
		n.Data = "em"
		return p.NewGeneric(n)
	case "h1", "h2", "h3", "h4", "h5":
		return p.NewHeading(n)
	case "strong", "b":
		n.Data = "strong"
		return p.NewGeneric(n)
	case "details":
		return p.parseMetadata(n)
	}

	return p.NewGeneric(n)
}
