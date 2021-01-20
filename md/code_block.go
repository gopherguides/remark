package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) NewCodeBlock(node *ast.CodeBlock) (*remark.CodeBlock, error) {
	cb := &remark.CodeBlock{
		Generic:  remark.NewGeneric("code"),
		Language: string(node.Info),
	}

	if len(cb.Language) == 0 {
		cb.Language = "plain"
	}

	return cb, nil
}
