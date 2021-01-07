package md

import (
	"bytes"
	"reflect"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) set(el remark.Element, node ast.Node) error {
	rv := reflect.ValueOf(el).Elem()
	f := rv.FieldByName("Elems")
	if f.IsValid() {
		elems, err := p.NewElements(node.GetChildren())
		if err != nil {
			return err
		}
		f.Set(reflect.ValueOf(elems))
	}

	if leaf := node.AsLeaf(); leaf != nil {
		at := leaf.Attribute
		if at != nil {
			atts := remark.Attributes{}
			for k, v := range at.Attrs {
				atts.Set(k, string(v))
			}
			atts.Set("id", string(at.ID))
			atts.Set("class", string(bytes.Join(at.Classes, []byte(" "))))
		}
	}

	return nil
}
