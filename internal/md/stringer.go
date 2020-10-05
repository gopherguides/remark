package md

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

func String(node ast.Node) string {
	switch t := node.(type) {
	case fmt.Stringer:
		return t.String()
	case *ast.Text:
		return string(t.Literal)
	case *ast.HorizontalRule:
		return string(t.Literal)
	case *ast.Code:
		return fmt.Sprintf("`%s`", string(t.Literal))
	case *ast.CodeBlock:
		return string(t.Literal)
	case *ast.HTMLSpan:
		return string(t.Literal)
	}
	return fmt.Sprintf("%T", node)

	// p := node.AsContainer()
	// if p == nil {
	// 	if p.AsLeaf() == nil {
	// 		return ""
	// 	}
	// 	// fmt.Printf(">>>TODO internal/md/stringer.go:13: node (%T): %v \n", node, node)
	// 	return ""
	// }
	//
	// bb := &bytes.Buffer{}
	//
	// for _, n := range p.GetChildren() {
	// 	l := n.AsLeaf()
	// 	if l == nil {
	// 		// fmt.Printf(">>>TODO internal/md/stringer.go:17: n (%T): %v \n", n, n)
	// 		continue
	// 	}
	//
	// 	if len(l.GetChildren()) > 0 {
	// 		bb.WriteString(String(l))
	// 	}
	// 	if len(l.Literal) > 0 {
	// 		bb.Write(l.Literal)
	// 	}
	// }
	// return bb.String()
}
