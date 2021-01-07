package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) NewImage(node *ast.Image) (remark.Tag, error) {
	image := &remark.Image{}
	err := p.set(image, node)
	return image, err
}
