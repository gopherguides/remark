package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewCodeBlock(node *ast.CodeBlock) (*remark.Code, error) {
	cb := &remark.Code{
		Language: string(node.Info),
		Elements: Elements{
			String(string(node.Literal)),
		},
	}

	if len(cb.Language) == 0 {
		cb.Language = "plain"
	}

	return cb, nil
}
