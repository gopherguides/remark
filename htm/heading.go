package htm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

var HEADS Headings

type Headings []string

func (h Headings) Allowed() []string {
	return []string{"h1", "h2", "h3", "h4", "h5"}
}

func (heads Headings) Valid(s string) bool {
	for _, h := range heads.Allowed() {
		if s == h {
			return true
		}
	}
	return false
}

func (p *Parser) NewHeading(n *html.Node) (*remark.Heading, error) {
	g, err := p.NewGeneric(n)
	if err != nil {
		return nil, err
	}

	if !HEADS.Valid(g.Name) {
		return nil, fmt.Errorf("invalid heading %v", n)
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
