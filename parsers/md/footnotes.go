package md

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) NewFootnotes(node *ast.Footnotes) (remark.Tag, error) {
	fmt.Printf(">>>TODO parsers/md/footnotes.go:11: node (%T): %v \n", node, node)
	return nil, fmt.Errorf("footnotes not supported")
}
