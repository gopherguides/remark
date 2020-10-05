package remark

import (
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

	// var section Section
	var fn ast.NodeVisitorFunc = func(node ast.Node, entering bool) ast.WalkStatus {
		if !entering {
			return ast.GoToNext
		}

		// if _, ok := node.(*ast.Document); !ok {
		// 	section.Nodes = append(section.Nodes, node)
		// }
		//
		// switch t := node.(type) {
		// case *ast.Document:
		// 	m.Node = t
		// case *ast.CodeBlock:
		// 	cb := CodeBlock{
		// 		CodeBlock: t,
		// 	}
		// 	section.Elements = append(section.Elements, cb)
		// case *ast.Paragraph:
		// 	p := Paragraph{
		// 		Paragraph: t,
		// 	}
		// 	section.Elements = append(section.Elements, p)
		// // 	section.Children = append(section.Children, &Paragraph{
		// // 		Paragraph: t,
		// // 	})
		// // 	return ast.SkipChildren
		// case *ast.Heading:
		// 	if !section.IsZero() {
		// 		m.Sections = append(m.Sections, section)
		// 	}
		// 	h := Heading{Heading: t}
		// 	section = Section{
		// 		Heading: h,
		// 	}
		// // 	section.Children = append(section.Children, &Heading{
		// // 		Heading: t,
		// // 	})
		// // 	return ast.SkipChildren
		// // case *ast.HorizontalRule:
		// // 	m.Sections = append(m.Sections, section)
		// // 	section = Section{}
		// default:
		// 	section.Nodes = append(section.Nodes, t)
		// 	// section.Elements = append(section.Elements, t)
		// 	fmt.Printf("%T\n", t)
		// }

		return ast.GoToNext
	}
	ast.Walk(node, fn)

	// if !section.IsZero() {
	// 	m.Sections = append(m.Sections, section)
	// }
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
