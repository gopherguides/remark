package htm

import (
	"strconv"
	"strings"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func (p *Parser) NewHeading(n *html.Node) (*remark.Heading, error) {
	g, err := p.NewGeneric(n)
	if err != nil {
		return nil, err
	}

	l := strings.TrimPrefix(g.Name, "h")
	lvl, err := strconv.Atoi(l)
	if err != nil {
		return nil, err
	}

	h := &remark.Heading{
		Generic: g,
		Level:   lvl,
	}
	return h, nil
}
