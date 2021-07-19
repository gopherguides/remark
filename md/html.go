package md

import (
	"fmt"
	"strings"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/htm"
	"golang.org/x/net/html"
)

func (p *Parser) parseHTML(src []byte) (remark.Tag, error) {
	s := strings.ToLower(string(src))
	if strings.HasPrefix(s, "<include") {
		doc, err := html.Parse(strings.NewReader(s))
		if err != nil {
			return nil, err
		}
		body, err := htm.HTMLBody(doc)
		if err != nil {
			return nil, err
		}
	}

	hp := htm.NewParser(p.Root)
	tag, err := hp.Parse(src)
	if err != nil {
		return nil, err
	}

	return tag, nil
}
