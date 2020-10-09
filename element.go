package remark

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gomarkdown/markdown/ast"
)

type Element interface {
	fmt.Stringer
	json.Marshaler
}

type Elements []Element

func (elems Elements) String() string {
	lines := make([]string, 0, len(elems))
	for _, e := range elems {
		lines = append(lines, e.String())
	}
	return strings.Join(lines, "")
}

func NewElements(nodes []ast.Node) (Elements, error) {
	elems := make(Elements, 0, len(nodes))
	for _, n := range nodes {
		el, err := NewElement(n)
		if err != nil {
			return nil, err
		}
		elems = append(elems, el)
	}
	return elems, nil
}

func NewElement(node ast.Node) (Element, error) {
	switch t := node.(type) {
	case *ast.Document:
		return NewModule(t)
	case *ast.Paragraph:
		return NewParagraph(t)
	case *ast.CodeBlock:
		return NewCodeBlock(t)
	case *ast.Heading:
		return NewHeading(t)
	case *ast.Text:
		return NewText(t)
	case *ast.HTMLSpan:
		return NewHTML(t)
	case *ast.Code:
		return NewCode(t)
	case *ast.HorizontalRule:
		return NewHR(t)
	case *ast.List:
		return NewList(t)
	case *ast.Link:
		return NewLink(t)
	case *ast.ListItem:
		return NewListItem(t)
	case *ast.Strong:
		return NewStrong(t)
	default:
		return nil, fmt.Errorf("unknown node type %T", node)
	}
	// return nil, nil
}
