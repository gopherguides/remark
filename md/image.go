package md

import (
	"github.com/gopherguides/remark"
	"github.com/markbates/markdownr/ast"
)

func (p *Parser) NewImage(i *ast.Image) (*remark.Generic, error) {
	g := remark.NewGeneric("image")
	g.Attributes["src"] = string(i.Destination)
	return g, nil
}
