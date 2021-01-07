package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)

func (p *Parser) NewOrderedList(n *html.Node) (Element, error) {
	ol := &remark.OrderedList{}

	if err := p.set(ol, n); err != nil {
		return nil, err
	}

	return ol, nil
}
