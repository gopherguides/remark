package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) NewImage(i *ast.Image) (*remark.Generic, error) {
	g := remark.NewGeneric("image")
	g.Attributes["src"] = string(i.Destination)
	return g, nil
}
