package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewImage(node *ast.Image) (*remark.Image, error) {
	image := &remark.Image{}
	err := set(image, node)
	return image, err
}
