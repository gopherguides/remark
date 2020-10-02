package remark

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/gopherguides/remark/internal/md"
)

type Parser struct {
	*parser.Parser
}

func (p *Parser) Parse(src []byte) (*Module, error) {
	if p.Parser == nil {
		p.Parser = md.NewParser()
	}
	node := markdown.Parse(src, p.Parser)

	m := &Module{}

	section := Section{}
	var fn ast.NodeVisitorFunc = func(node ast.Node, entering bool) ast.WalkStatus {
		if !entering {
			return ast.GoToNext
		}
		switch t := node.(type) {
		case *ast.Document:
			m.Node = t
		case *ast.Paragraph:
			section.Children = append(section.Children, &Paragraph{
				Paragraph: t,
			})
			return ast.SkipChildren
		case *ast.Heading:
			section.Children = append(section.Children, &Heading{
				Heading: t,
			})
			return ast.SkipChildren
		case *ast.HorizontalRule:
			m.Sections = append(m.Sections, section)
			section = Section{}
		default:
			fmt.Printf("%T\n", node)
		}
		// switch t := node.(type) {
		// case *ast.HorizontalRule:
		// 	section.Children = append(section.Children, t)
		// 	m.Sections = append(m.Sections, section)
		// 	section = Section{}
		// 	return ast.GoToNext
		// case *ast.Heading:
		// 	h := &Heading{
		// 		Heading: t,
		// 	}
		// 	section.Children = append(section.Children, h)
		// 	// case *ast.Text:
		// 	// fmt.Println(string(t.Literal))
		// case *ast.Paragraph:
		// 	section.Children = append(section.Children, t)
		// 	return ast.SkipChildren
		// default:
		// 	// section.Children = append(section.Children, t)
		// }

		return ast.GoToNext
	}
	ast.Walk(node, fn)

	return m, nil
}

func (p *Parser) ParseFile(filename string) (*Module, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return p.Parse(b)
}

func (p *Parser) ParseReader(r io.Reader) (*Module, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return p.Parse(b)
}

func Parse(src []byte) (*Module, error) {
	p := &Parser{}
	return p.Parse(src)
}

func ParseFile(filename string) (*Module, error) {
	p := &Parser{}
	return p.ParseFile(filename)
}

func ParseReader(r io.Reader) (*Module, error) {
	p := &Parser{}
	return p.ParseReader(r)
}
