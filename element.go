package remark

import (
	"github.com/gomarkdown/markdown/ast"
)

type Element interface {
	// fmt.Stringer
}

type Elements []Element

// func (elems Elements) String() string {
// 	lines := make([]string, 0, len(elems))
// 	for _, e := range elems {
// 		lines = append(lines, e.String())
// 	}
// 	return strings.Join(lines, "\n\n")
// }

func NewElement(node ast.Node) (Element, error) {
	switch t := node.(type) {
	case *ast.Document:
		return NewModule(t)
	case *ast.Paragraph:
		return NewParagraph(t)
	case *ast.CodeBlock:
		return NewCodeBlock(t)
	}
	return nil, nil
}
