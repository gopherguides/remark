package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)

func (p *Parser) NewTable(n *html.Node) (Element, error) {
	table := &remark.Table{}

	err := p.set(table, n)
	return table, err
}
