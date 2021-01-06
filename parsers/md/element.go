package md

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/parsers/htm"
)

type Element = remark.Element
type Elements = remark.Elements

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
		return NewDocument(t)
	case *ast.Paragraph:
		return NewParagraph(t)
	case *ast.CodeBlock:
		return NewCodeBlock(t)
	case *ast.Heading:
		return NewHeading(t)
	case *ast.Text:
		return remark.String(string(t.Literal)), nil
	case *ast.HTMLSpan:
		return htm.Parse(t.Literal)
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
	case *ast.Emph:
		return NewEmph(t)
	case *ast.Del:
		return NewDel(t)
	case *ast.BlockQuote:
		return NewBlockQuote(t)
	case *ast.Table:
		return NewTable(t)
	case *ast.TableHeader:
		return NewTableHeader(t)
	case *ast.TableRow:
		return NewTableRow(t)
	case *ast.TableCell:
		return NewTableCell(t)
	case *ast.TableBody:
		return NewTableBody(t)
	case *ast.Image:
		return NewImage(t)
	case *ast.Superscript:
		return NewSuperScript(t)
	case *ast.Subscript:
		return NewSubScript(t)
	case *ast.Footnotes:
		return NewFootnotes(t)
	case *ast.HTMLBlock:
		return htm.Parse(t.Literal)
	}
	return nil, fmt.Errorf("unknown node type %T", node)
}
