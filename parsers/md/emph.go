package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewEmph(node *ast.Emph) (*remark.Emph, error) {
	s := &remark.Emph{}

	for _, n := range node.Children {
		el, err := NewElement(n)
		if err != nil {
			return nil, err
		}
		s.Elements = append(s.Elements, el)
	}
	return s, nil
}
