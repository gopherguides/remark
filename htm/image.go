package htm

import (
	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func (p *Parser) NewImage(n *html.Node) (*remark.Image, error) {
	g, err := p.NewGeneric(n)
	if err != nil {
		return nil, err
	}

	src := g.Attributes["src"]
	inf, err := p.NewInfo(p.Info.Join(src))
	if err != nil {
		return nil, err
	}

	inc := &remark.Image{
		Info:    inf,
		Generic: g,
	}

	return inc, nil
}
