package htm

import (
	"fmt"
	"strings"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func (p *Parser) NewGeneric(n *html.Node) (remark.Tag, error) {
	t := &remark.Generic{
		TagName: strings.ToLower(n.Data),
	}
	if err := p.set(t, n); err != nil {
		return nil, err
	}

	if c := n.FirstChild; c != nil {
		fmt.Printf(">>>TODO parsers/htm/generic.go:20: c (%T): %v \n", c, c)
		el, err := p.NewElement(c)
		if err != nil {
			return nil, err
		}
		fmt.Printf(">>>TODO parsers/htm/generic.go:25: el (%T): %v \n", el, el)
	}
	return t, nil
}
