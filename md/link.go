package md

import (
	"github.com/gopherguides/remark"
	"github.com/markbates/markdownr/ast"
)

func (p *Parser) parseLink(link *ast.Link) (remark.Tag, error) {
	tag := remark.NewGeneric("a")

	tag.Set("href", string(link.Destination))

	t := string(link.Title)
	if len(t) > 0 {
		tag.Set("title", t)
	}

	for _, n := range link.Children {
		nt, err := p.parseNode(n)
		if err != nil {
			return nil, err
		}
		if nt != nil {
			tag.Children = append(tag.Children, nt)
		}
	}

	return tag, nil
}
