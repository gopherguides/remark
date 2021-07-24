package htm

import (
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
