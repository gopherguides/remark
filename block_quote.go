package remark

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type BlockQuote struct {
	*ast.BlockQuote
	Elements Elements
}

func (bq BlockQuote) String() string {
	return fmt.Sprintf("> %s", bq.Elements)
}

func NewBlockQuote(node *ast.BlockQuote) (*BlockQuote, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	bq := &BlockQuote{
		BlockQuote: node,
		Elements:   elems,
	}

	return bq, nil
}
