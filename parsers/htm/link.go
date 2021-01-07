package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)

func (p *Parser) NewLink(n *html.Node) (Element, error) {
	a := &remark.Link{}
	err := p.set(a, n)
	return a, err
}
