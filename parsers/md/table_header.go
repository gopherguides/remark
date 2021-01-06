package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewTableHeader(node *ast.TableHeader) (*remark.TH, error) {
	th := &remark.TH{}
	err := set(th, node)
	return th, err
}
