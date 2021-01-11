package md

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

type Element = remark.Element
type Elements = remark.Elements

func (p *Parser) NewElements(nodes []ast.Node) (Elements, error) {
	elems := make(Elements, 0, len(nodes))
	for _, n := range nodes {
		el, err := p.NewElement(n)
		if err != nil {
			return nil, err
		}
		elems = append(elems, el)
	}
	return elems, nil
}

func (p *Parser) NewElement(node ast.Node) (Element, error) {
	switch t := node.(type) {
	case *ast.Document:
		return p.NewDocument(t)
	case *ast.Paragraph:
		return p.NewGeneric("p", t)
	case *ast.CodeBlock:
		return p.NewCodeBlock(t)
	case *ast.Heading:
		return p.NewHeading(t)
	case *ast.Text:
		return remark.String(string(t.Literal)), nil
	case *ast.Code:
		return p.NewGeneric("code", t)
	case *ast.HorizontalRule:
		return NewHR()
	case *ast.List:
		return p.NewList(t)
	case *ast.Link:
		return p.NewLink(t)
	case *ast.ListItem:
		return p.NewGeneric("li", t)
	case *ast.Strong:
		return p.NewGeneric("strong", t)
	case *ast.Emph:
		return p.NewGeneric("em", t)
	case *ast.Del:
		return p.NewGeneric("strike", t)
	case *ast.BlockQuote:
		return p.NewGeneric("blockquote", t)
	case *ast.Table:
		return p.NewTable(t)
	case *ast.TableHeader:
		return p.NewGeneric("thead", t)
	case *ast.TableRow:
		return p.NewGeneric("tr", t)
	case *ast.TableCell:
		return p.NewGeneric("td", t)
	case *ast.TableBody:
		return p.NewGeneric("tbody", t)
	case *ast.Image:
		return p.NewImage(t)
	case *ast.Superscript:
		return p.NewGeneric("sup", t)
	case *ast.Subscript:
		return p.NewGeneric("sub", t)
	case *ast.Footnotes:
		return p.NewFootnotes(t)
	case *ast.HTMLSpan:
		return p.htmlBody(t.Literal)
	case *ast.HTMLBlock:
		return p.htmlBody(t.Literal)
	}
	return nil, fmt.Errorf("unknown node type %T", node)
}
