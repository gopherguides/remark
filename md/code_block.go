package md

import (
	"github.com/gopherguides/remark"
	"github.com/markbates/markdownr/ast"
)

func (p *Parser) NewCode(node *ast.Code) (*remark.Code, error) {
	g := remark.NewGeneric("code")

	g.Children = append(g.Children, remark.String(string(node.Literal)))

	cb := &remark.Code{
		Generic: g,
	}

	return cb, nil
}

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

	return cb, nil
}
