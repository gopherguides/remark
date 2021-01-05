package md

import (
	"html/template"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark/parsers/htm"
)

type HTMLBlock struct {
	*ast.HTMLBlock
	Elements Elements
}

func (hb HTMLBlock) String() string {
	return hb.Elements.String()
}

func (hb HTMLBlock) HTML() template.HTML {
	return hb.Elements.HTML()
}

func NewHTMLBlock(node *ast.HTMLBlock) (*HTMLBlock, error) {
	hb := &HTMLBlock{
		HTMLBlock: node,
	}

	b := node.Leaf.Literal

	els, err := htm.Parse(b)
	if err != nil {
		return nil, err
	}
	hb.Elements = append(hb.Elements, els)
	return hb, nil
}
