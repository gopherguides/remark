package remark

import "github.com/gomarkdown/markdown/ast"

type Text struct {
	*ast.Text
}

func NewText(node *ast.Text) (Text, error) {
	t := Text{
		Text: node,
	}
	return t, nil
}

func (t Text) String() string {
	return string(t.Literal)
}
