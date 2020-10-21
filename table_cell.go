package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

// TableCell is strikethrough
type TableCell struct {
	*ast.TableCell
	Elements Elements
}

func (cell TableCell) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"content": cell.Elements,
		"literal": cell.String(),
	}

	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", cell): m,
	}, "", "  ")
}

func (cell TableCell) String() string {
	s := cell.Elements.String()
	return s
}

func NewTableCell(node *ast.TableCell) (*TableCell, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	cell := &TableCell{
		TableCell: node,
		Elements:  elems,
	}
	return cell, nil
}
