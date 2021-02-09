package htm

import (
	"fmt"
	"strings"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func (p *Parser) NewGeneric(n *html.Node) (*remark.Generic, error) {
	g := remark.NewGeneric(n.Data)

	for _, at := range n.Attr {
		g.Set(at.Key, at.Val)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tg, err := p.NewTag(c)
		if err != nil {
			return nil, err
		}
		if tg != nil {
			g.Children = append(g.Children, tg)
		}
	}

	return g, nil
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
	case "hr":
		return p.NewGeneric(n)
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
