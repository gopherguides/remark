package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewBlockQuote(node *ast.BlockQuote) (*remark.BlockQuote, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	bq := &remark.BlockQuote{
		Elements: elems,
	}

	return bq, nil
}
