package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type BlockQuote struct {
	*ast.BlockQuote
	Elements Elements
}

func (bq BlockQuote) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"content": bq.Elements,
		"literal": bq.String(),
	}

	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", bq): m,
	}, "", "  ")
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
