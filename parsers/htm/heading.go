package htm

import (
	"strconv"
	"strings"

	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)

func (p *Parser) NewHeading(n *html.Node) (Element, error) {
	h := &remark.Heading{}

	l := strings.TrimPrefix("h", n.Data)
	lvl, err := strconv.Atoi(l)
	if err != nil {
		return nil, err
	}
	h.Level = lvl

	err = p.set(h, n)
	return h, err
}
