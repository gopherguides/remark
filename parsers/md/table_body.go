package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewTableBody(node *ast.TableBody) (*remark.TableBody, error) {
	body := &remark.TableBody{}
	err := set(body, node)
	return body, err
}
