package htm

import (
	"strings"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func (p *Parser) NewGeneric(n *html.Node) (remark.Tag, error) {
	t := &remark.Generic{
		TagName: strings.ToLower(n.Data),
	}
	err := p.set(t, n)
	return t, err
}
