package md

import (
	"fmt"

	"github.com/gopherguides/remark"

	"github.com/markbates/markdownr/ast"
)

func (p *Parser) NewHeading(node *ast.Heading) (*remark.Heading, error) {
	tag := &remark.Heading{
		Generic: remark.NewGeneric(fmt.Sprintf("h%d", node.Level)),
		Level:   node.Level,
	}

	for _, n := range node.Children {
		t, err := p.parseNode(n)
		if err != nil {
			return nil, err
		}
		tag.Append(t)
	}
	return tag, nil
}
