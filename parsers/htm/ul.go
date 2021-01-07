package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)

func (p *Parser) NewUnorderedList(n *html.Node) (Element, error) {
	ul := &remark.UnorderedList{}

	err := p.set(ul, n)
	return ul, err
}
