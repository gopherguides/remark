package md

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopherguides/remark"
	"github.com/markbates/markdownr/ast"
)

func (p *Parser) NewImage(i *ast.Image) (*remark.Generic, error) {
	g := remark.NewGeneric("image")
	panic(g)
	src := string(i.Destination)
	g.Attributes["src"] = src

	fp := filepath.Join(p.Root, src)
	fmt.Println("TODO >> fp", fp)
	if _, err := os.Stat(fp); err != nil {
		return nil, err
	}
	return g, nil
}
