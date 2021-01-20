package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) NewCodeBlock(node *ast.CodeBlock) (*remark.CodeBlock, error) {
	g := remark.NewGeneric("code")

	g.Children = append(g.Children, remark.String(string(node.Literal)))

	cb := &remark.CodeBlock{
		Generic:  g,
		Language: string(node.Info),
	}

	if len(cb.Language) == 0 {
		cb.Language = "plain"
	}

	// for _, n := range node.Children {
	// 	nt, err := p.parseNode(n)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if nt != nil {
	// 		t.Children = append(t.Children, nt)
	// 	}
	// }
	return cb, nil
}
