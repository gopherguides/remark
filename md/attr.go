package md

import (
	"bytes"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func attrs(at *ast.Attribute) remark.Attributes {
	atts := remark.Attributes{}
	if at == nil {
		return atts
	}

	for k, v := range a.Attrs {
		atts[k] = string(v)
	}
	atts["id"] = string(a.ID)
	atts["class"] = string(bytes.Join(a.Classes, []byte(" ")))
	return atts
}
